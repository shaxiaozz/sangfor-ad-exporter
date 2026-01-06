package sangfor_ad

import (
	"fmt"
	"github.com/shaxiaozz/sangfor-ad-exporter/global"
	"github.com/shaxiaozz/sangfor-ad-exporter/model"
)

// 登录获取Token
func Login(req *model.SangforAdLoginReq) (*model.SangforAdLoginResp, error) {
	// 封装请求体
	reqUrl := fmt.Sprintf("%s/api/token", global.Config.SangforAd.Url)

	return requestPost[model.SangforAdLoginResp](reqUrl, "POST", req, nil)
}

// 获取虚拟服务状态信息
func VirtualServiceStat(token string) (*model.SangforAdVirtualServiceStatResp, error) {
	// 封装请求体
	reqUrl := fmt.Sprintf("%s/api/ad/v3/stat/slb/virtual-service", global.Config.SangforAd.Url)

	headers := map[string]string{
		"x-token-sangforad": token,
	}

	return requestGet[model.SangforAdVirtualServiceStatResp](reqUrl, headers)
}
