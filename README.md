#### go log system

<p align='left'>
<img src="https://img.shields.io/badge/build-passing-brightgreen.svg">
<a href="https://twitter.com/perfactsen"><img src="https://img.shields.io/badge/twitter-keke-green.svg?style=flat&colorA=009df2"></a>
<a href="https://www.zhihu.com/people/sencoed.com/activities"><img src="https://img.shields.io/badge/%E7%9F%A5%E4%B9%8E-keke-green.svg?style=flat&colorA=009df2"></a>
</p>

log is a powerful logging framework that provides log custom log level.

log provides Fatal, Error ,  Warn, Info, Debug level log. 



#### How to Use

Use Log just as you would use print.

```go
func main(){
	ctx := context.Background()
	
	v, err := rand.Int(rand.Reader, big.NewInt(int64(16)))
    	if err != nil {
    		log.ErrorContext(ctx, " rand.In failed", "error", err.Error())
    		return 
    	}
	log.InfoContext(ctx, "the rand Int result", "v", v)
}
```

This use is made of the key and the Value output.


### License

This is free software distributed under the terms of the MIT license
