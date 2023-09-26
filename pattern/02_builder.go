package pattern

/*
Реализовать паттерн "Строитель", объяснить применимость паттерна, плюсы и минусы, а
также реальные примеры его использования на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
Строитель - это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Применимость:
1) когда код должен создавать разные представления какого-то объекта;
2) когда нужно собирать сложные составные объекты.

Плюсы:
- позволяет создавать объекты пошагово;
- позволяет переиспользовать один и тот же код для создания различных объектов;
- изолирует сложный код сборки объекта от его основной бизнес-логики.

Минусы:
- усложняет код программы за счет дополнительных классов.
- клиент будет привязан к конкретным классам строителей, так как в интерфейсе строителя может не быть метода получения
результата

Примеры:
1) В Яндекс Еде почти в любом ресторане есть возможность выбрать из чего будет состоять блюдо. Т.е. мы "строим" наш условный
бургер из выбранного типа булочки, котлеты, а также из выбранных дополнительных ингредиентов.
2) Производство почти любой техники. Например, модели ноутбуков выходят с большим количеством вариаций комплектующих.
Т.е. в данном случае, строителем будет конвейерная лента или завод по производству ноутбуков.

Реализация:
Создание http запроса. Мы можем создать запрос на наш "Сервер 1" или же полностью самостоятельно его настроить.
*/

// HttpRequestBuilder интерфейс пошагового создания условного http запроса
type HttpRequestBuilder interface {
	SetURL(url string)
	SetHeaders(headers []string)
	SetParams(params map[string]string)
	SetMethod(method string)
	SetBody(body []byte)
	SetCookie(cookie []string)
}

type Request struct {
	url     string
	headers []string
	params  map[string]string
	method  string
	body    []byte
	cookie  []string
}

func (r Request) SetURL(url string) {
	r.url = url
}

func (r Request) SetHeaders(headers []string) {
	r.headers = headers
}

func (r Request) SetParams(params map[string]string) {
	r.params = params
}

func (r Request) SetMethod(method string) {
	// имитируем проверку на поданного метода
	if method != "GET" && method != "POST" {
		r.method = "GET"
	}
	r.method = method
}

func (r Request) SetBody(body []byte) {
	r.body = body
}

func (r Request) SetCookie(cookie []string) {
	r.cookie = cookie
}

// httpRequestDirector либо строит запросы к "Серверу 1", либо дает пользователю полностью настроить запрос
type httpRequestDirector struct {
	reqBuilder *HttpRequestBuilder
}

// GetServerOne посылает Get запрос на условный сервер 1
func (d *httpRequestDirector) GetServerOne() {
	(*d.reqBuilder).SetURL("https://serveroneurl")
	(*d.reqBuilder).SetMethod("GET")
}

// PostServerOne посылает Post запрос на условный сервер 1
func (d *httpRequestDirector) PostServerOne() {
	(*d.reqBuilder).SetURL("https://serveroneurl")
	(*d.reqBuilder).SetMethod("POST")
	(*d.reqBuilder).SetBody([]byte("Hello, server 1!"))
}

// CustomRequest посылаем полностью настроенный пользователем запрос
func (d *httpRequestDirector) CustomRequest(url string, headers []string, params map[string]string, method string, body []byte, cookie []string) {
	(*d.reqBuilder).SetURL(url)
	(*d.reqBuilder).SetHeaders(headers)
	(*d.reqBuilder).SetParams(params)
	(*d.reqBuilder).SetMethod(method)
	(*d.reqBuilder).SetBody(body)
	(*d.reqBuilder).SetCookie(cookie)
}
