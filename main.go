package main

import (
	"BiliMallRepublish/bilibili"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {
	// 程序结束后等待10s再退出
	defer func() {
		fmt.Println("程序运行结束，10秒后将自动退出...")
		time.Sleep(10 * time.Second)
	}()

	// 用户sessdata输入
	sessData := ""
	for {
		fmt.Print("请输入您的SessData:")
		i, err := fmt.Scanf("%s\n", &sessData)
		if i > 0 && err == nil {
			break
		}
	}
	sessData = strings.TrimSuffix(sessData, "\n")

	// 设置sessdata
	bilibili.SetSessdata(sessData)

	// 获得账号内市集已上架商品列表
	err, items := bilibili.GetNowPublishedList()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 无上架商品或sessdata无效
	if len(items) == 0 {
		fmt.Println("未发现账号市集内有效的已上架商品。")
		return
	}

	// 循环所有市集内已上架商品
	for _, item := range items {
		fmt.Print(fmt.Sprintf("正在下架 <%s> ...", item.C2CItemsName))

		// 通过C2CItemsId下架商品
		err = bilibili.DropItem(item.C2CItemsId)
		if err != nil {
			fmt.Println("下架商品发生致命错误！跳过。[" + err.Error() + "]")
			continue
		}
		fmt.Println("成功！")

		// 将BlindBoxId收集起来准备重新上架
		fmt.Println(fmt.Sprintf("准备重新上架 <%s> ...", item.C2CItemsName))
		itemsList := make([]string, 0)
		for _, detailDtoItem := range item.DetailDtoList {
			fmt.Println(fmt.Sprintf("捡起物件 <%s>", detailDtoItem.Name))
			itemsList = append(itemsList, strconv.Itoa(detailDtoItem.BlindBoxId))
		}

		// 延迟1秒
		//time.Sleep(time.Second)

		// call 市集 check接口，将收集的BlindBoxId一起传过去，来获取后续操作的token
		fmt.Print(fmt.Sprintf("对捡起的%d物件进行准备...", len(itemsList)))
		err, token, showTime := bilibili.CheckItems(itemsList)
		if err != nil {
			fmt.Println("发生致命错误！跳过。[" + err.Error() + "]")
			continue
		}
		fmt.Println("成功！")
		fmt.Print(fmt.Sprintf("<%s> ，售价：%s ，剩余时间：%s", item.C2CItemsName, item.ShowPrice, showTime))

		// 延迟1秒
		//time.Sleep(time.Second)

		// call publish接口，先将isConfirm设置为false请求一次（其实应该可以跳过这步，目前是为了被检测作为保险）
		err, _, discount := bilibili.PublishItem(item.ShowPrice, token, false)
		if err != nil {
			fmt.Println("...发生致命错误！跳过。[" + err.Error() + "]")
			continue
		}
		fmt.Println(fmt.Sprintf("，商品折扣：%s", discount))
		//time.Sleep(time.Second)
		fmt.Print(fmt.Sprintf("开始执行上架商品 <%s> ...", item.C2CItemsName))

		// 第二次call publish接口将isConfirm设置为true正式上架
		err, isPublish, _ := bilibili.PublishItem(item.ShowPrice, token, true)
		if err != nil {
			fmt.Println("发生致命错误！跳过。[" + err.Error() + "]")
			continue
		}
		if isPublish {
			fmt.Println("成功！")
		} else {
			fmt.Println("失败！")
		}

		// 一件商品上下架完成后延迟1秒，循环下一件
		//time.Sleep(time.Second)
	}
}
