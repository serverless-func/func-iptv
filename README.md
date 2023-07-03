## fun-iptv

### 说明

数据来源 https://github.com/BurningC4/Chinese-IPTV
用于加速 m3u 以及 epg 的访问

### IPTV LIST

| Name        | Url                                  |
| ----------- | ------------------------------------ |
| CCTV 频道   | https://iptv.fun.dongfg.com/cctv.m3u |
| CCTV 节目单 | https://iptv.fun.dongfg.com/cctv.xml |

### Deployment

## Development

```shell
fission -n fission spec init
fission -n fission fn create --spec --name func-iptv --src src.zip --entrypoint Handler --env go --buildcmd "./customBuild.sh"
fission -n fission route create --spec --method GET --name func-iptv --url /{Subpath} --function func-iptv --createingress  --ingressrule "iptv.func.dongfg.com=/" --ingresstls "tls-iptv-func-dongfg" --ingressannotation "cert-manager.io/cluster-issuer=letsencrypt-dongfg"
```
