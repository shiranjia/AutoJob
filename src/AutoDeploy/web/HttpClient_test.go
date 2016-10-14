package web

import (
	"testing"
)

func TestHttpDial(t *testing.T)  {
	head := make(map[string]string)
	head["Content-Type"] = "application/json;charset=UTF-8"
	//head["Content-Type"]="application/x-www-form-urlencoded"
	args := make(map[string]string)
	args["cityId"] = "36"
	res := httpDial("http://m-piao.jd.com/rest/search","POST",head,args)
	t.Log(res)
	if res == "" {
		t.Error("res is null")
	}

}

func TestHttpDialBody(t *testing.T)  {
	head := make(map[string]string)
	head["Content-Type"] = "application/xml;charset=UTF-8"
	testURL := "http://gw.piao.jd.com:171/"
	proURL := "http://gw.piao.360buy.com/"
	//updateProduct.action getCityInfo.action
	localURL := "http://gw.piao.jd.com/"
	cityRequest := "<?xml version=\"1.0\" encoding=\"UTF-8\" ?><request> <head> <version> 1.0</version> <msgId>1445936150237</msgId> <partnerCode>10000</partnerCode> <timeStamp>20151027165550237</timeStamp> <signed>YjA5ZmMyYjIwMDBmNmI1Ng==</signed> </head> <body> <getCityInfo> <cityId>36</cityId> </getCityInfo> </body> </request>"
	//a :="<?xml version=\"1.0\" encoding=\"UTF-8\"?><request><head><version>1.0</version><msgId>1468463654539</msgId><partnerCode>10000</partnerCode><timeStamp>20160714103414540</timeStamp><signed>ZTZkN2YxZDk4ZDJmNGRmZg==</signed></head><body><updateProduct><product><id>87665888</id><name><![CDATA[赖声川导演监制 亲子音乐剧《流浪狗之歌》]]></name><detail><![CDATA[<p style=\"text-indent:2em;\">艺术总监 | 赖声川</p><p style=\"text-indent:2em;\">编剧/导演 | 赵自强</p><p style=\"text-indent:2em;\">音乐总监 | 洪予彤</p><p style=\"text-indent:2em;\">主演 | 鲍春来、罗永娟、阮力等</p><p style=\"text-indent:2em;\"><strong><span style=\"color:#ff6600;\">*具体信息以现场演出为准 </span></strong></p><p style=\"text-indent:2em;\"><strong><span style=\"color:#ff6600;\"><br /></span></strong></p><p style=\"text-align:center\"><img src=\"http://static.228.cn/upload/Image/201602/1455852534111_9126_x.jpg\" alt=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" width=\"450\" height=\"301\" border=\"0\" hspace=\"0\" vspace=\"0\" title=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" style=\"width:450px;height:301px;\" /></p><p style=\"text-align:center\"><br /></p><p style=\"text-indent:2em;\"><strong>剧情介绍</strong></p><p style=\"text-indent:2em;\">Lucky是一只顽皮的宠物狗，不小心迷失了回家的方向。第一次离开家的它，既害怕又伤心，不知道自己是否能找到回家的路。Lucky在返家的路途中，遇见了一群流浪狗：有生过很多小狗的老妈妈皇后、流浪街头多年的少女狗宝贝、生了皮肤病的癞皮狗毛毛、有过五个主人的老师，还有被主人丢掉三次的路克。初次流浪的Lucky要怎么和它们生活？可怕的捕狗人还在四处搜索，想把这群流浪狗统统抓走，Lucky能圆回家的梦吗？流浪的生命能否有出路呢？《流浪狗之歌》透过歌舞剧的形式，要与亲子一起分享负责任的爱，以及对生命的尊重。</p><p style=\"text-indent:2em;\"><br /></p><p style=\"text-align:center\"><img src=\"http://static.228.cn/upload/Image/201602/1455852560657_640_x.jpg\" alt=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" width=\"450\" height=\"537\" border=\"0\" hspace=\"0\" vspace=\"0\" title=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" style=\"width:450px;height:537px;\" /></p><p style=\"text-align:center\"><br /></p><p style=\"text-indent:2em;\"><strong>服装造型</strong></p><p style=\"text-indent:2em;\">音乐剧中，六个主要角色分别有不同造型，Lucky是雪白的比熊犬，路克是帅气的哈士奇，而Baby是可爱的米克斯，在造型上皆有些许的象征。皇后是高贵的马尔济斯，老师是憨直的松狮犬，而毛毛则是癞皮狗。</p><p style=\"text-indent:2em;\"><br /></p><p style=\"text-align:center\"><img src=\"http://static.228.cn/upload/Image/201602/1455852575988_1977_x.jpg\" alt=\"赖声川导演监制 亲子音乐剧《流浪狗之歌\" width=\"450\" height=\"301\" border=\"0\" hspace=\"0\" vspace=\"0\" title=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" style=\"width:450px;height:301px;\" /></p><p style=\"text-align:center\"><br /></p><p style=\"text-indent:2em;\"><strong>主要演员：</strong></p><p style=\"text-indent:2em;\"><strong>鲍春来</strong></p><p style=\"text-indent:2em;\">曾是中国国家羽球队运动员。2006年韩国公开赛单打冠军，2007年中国羽毛球超级赛冠军。拥有亚锦赛、德国公开赛、日本超级赛、韩国公开赛、新加坡超级赛等一系列单打冠军头衔。并曾是2005、2011苏迪曼杯和2004、2006、2008、2010汤姆斯杯团体冠军成员。2011年退役后，鲍春来事业重心转到演艺圈，成为一名演员。曾在旅游探险类真人秀节目《我是冒险王》中担任主持；曾以嘉宾身份参加过《天天向上》《快乐大本营》《我不是明星》等娱乐类节目。2015年参与丁乃筝舞台剧《他和他的两个老婆》大陆巡回演出。在《流浪狗之歌》中担任主演路克。</p><p style=\"text-indent:2em;\"><br /></p><p style=\"text-align:center\"><img src=\"http://static.228.cn/upload/Image/201602/1455852591229_5656_x.jpg\" alt=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" width=\"450\" height=\"301\" border=\"0\" hspace=\"0\" vspace=\"0\" title=\"赖声川导演监制 亲子音乐剧《流浪狗之歌》\" style=\"width:450px;height:301px;\" /></p><p style=\"text-align:center\"><br /></p><p style=\"text-indent:2em;\"><strong>罗永娟</strong></p><p style=\"text-indent:2em;\">歌手、演员，CCTV梦想中国女子组冠军，歌曲代表作有：电视剧《聊斋》片尾曲《胭脂泪》、《加林赛部落》、《巧克力女孩》等，影视代表作有：新版《西游记》（饰哪吒），电视剧《十二生肖传奇》（饰演巴托）、赖声川舞台剧《十三角关系》（饰安琪）、音乐剧《如果有来生》。</p><p style=\"text-indent:2em;\"><br /></p><p style=\"text-indent:2em;\"><strong>阮力</strong></p><p style=\"text-indent:2em;\">歌手、演员，2008年华南赛区街舞大赛第一名，2010年参与湖南视《快乐男声》晋级迅雷赛区全国五强而广为人知，2012年《奋勇争先娱乐圈》年度亚军。舞台剧作品有2013赖声川导演的《如梦之梦》、2014赖声川导演的《海鸥》。</p><p style=\"text-indent:2em;\"><br /></p><p style=\"text-indent:2em;\">2016上海赖声川导演监制 亲子音乐剧《流浪狗之歌》，敬请期待！</p>]]></detail><info><![CDATA[艺术总监 | 赖声川编剧/导演 | 赵自强乐总监 | 洪予彤主演 | 鲍春来、罗永娟、阮力等*具体信息以现场演出为准 剧情介绍Lucky是一只顽皮的宠物狗，不小心迷失了回家的方向。第一次离开家的它，既害怕又伤心，不知道自己是否找到回家的路。Lucky在返家的路途中，遇见了一群流浪狗：有生过很多小狗的老妈妈皇后、流浪街头多年的少女狗宝贝、生了皮肤病的癞皮狗毛毛、有过五个主人的老师，还有被主人丢掉三次的克。初次流浪的Lucky要怎么和它们生活？可怕的捕狗人还在四处搜索，想把这群流浪狗统统抓走，Lucky能圆回家的梦吗？流浪的生命能否有出路呢？《流浪狗之歌》透过歌舞剧的形式，要与亲子一起分享负责任的爱，以及对生命的尊重。服装造型音乐剧中，六个主要角色分别有不同的造型，Lucky是雪白的比熊犬，路克是帅气的哈士奇，而Baby是可爱的米克斯，在造型上皆有些许的象征皇后是高贵的马尔济斯，老师是憨直的松狮犬，而毛毛则是癞皮狗。主要演员：鲍春来曾是中国国家羽毛球队运动员。2006年韩国公开赛单打冠军，2007年中国羽毛球超级赛冠军。拥有亚锦赛、德国公开赛、日本超级赛、韩国公开赛、新加坡超级赛...]]></info><offState>1</offState><sellState>0</sellState><cityId>108</cityId><cityName>上海</cityName><venueId>143718</venueId><venueName>东方艺术中心</venueName><typeId>36</typeId><typeName>儿童亲子</typeName><subtypeId>37</subtypeId><subtypeName>儿童剧</subtypeName><startTime>2016-07-16 00:00:00</startTime><endTime>2016-07-17 00:00:00</endTime><soleAgent>0</soleAgent><isSponsor>0</isSponsor></product></updateProduct></body></request>";
	_,_ ,_= testURL,proURL,localURL
	res := httpDialBody(testURL + "getCityInfo.action","POST",head,cityRequest)
	t.Log("\n",res)
	if res == "" {
		t.Error("res is null")
	}
}

func Benchmark_HttpDial(b *testing.B) {
	for i:=0;i<b.N ;i++{
		head := make(map[string]string)
		head["Content-Type"] = "application/json;charset=UTF-8"
		args := make(map[string]string)
		//res := httpDial("http://h5.ticket.jd.com//rest/cmsShop","GET",head,args)
		res := httpDial("http://h5.ticket.jd.com/rest/allCity","GET",head,args)
		_ = res
		//b.Log(res)
	}
}

//go test  -test.bench=".*"
//go test  -test.bench="Benchmark_HttpDial"