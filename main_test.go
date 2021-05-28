package main_test

import (
	"fmt"
	"git.900sui.cn/kc/kcgin/logs"
	"strconv"
	"testing"
)

func TestCm(t *testing.T) {
	//utils.CreateModel("shop_single")
	var TailorSubHead = []string{
		40:  "我要掏耳朵",
		41:  "我要牙齿美白",
		42:  "我要牙齿矫正",
		43:  "我要修补龋齿",
		44:  "我要穿耳洞",
		45:  "我要头部放松",
		46:  "我要不头胀",
		47:  "我要不失眠",
		48:  "我要不多梦",
		49:  "我要不头痛",
		50:  "我要不头晕",
		51:  "我要去头屑",
		52:  "我要祛白发",
		53:  "我要治秃顶",
		54:  "我要治脱发",
		55:  "我要植发",
		56:  "我要染发",
		57:  "我要接发",
		58:  "我要洗剪吹",
		59:  "我要烫发",
		60:  "我要纹身",
		92:  "我要护发",
		138: "我要洗头",
		148: "我要吹发",
		149: "我要盘发",
		150: "我要洗眉毛",
	}

	for k, v := range TailorSubHead {
		if len(v) == 0 {
			continue
		}
		logs.Info(k, v)
	}

}

func TestT(t *testing.T) {
	parseInt, _ := strconv.ParseInt("1596676386", 10, 64)
	fmt.Println(parseInt)
}
