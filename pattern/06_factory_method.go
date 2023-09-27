package pattern

/*
Реализовать паттерн "Фабричный метод", объяснить применимость паттерна, плюсы и минусы, а
также реальные примеры его использования на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern

Фабричный метод - это порождающий паттерн проектирования, который определяет общий интерфейс для создания объектов в
суперклассе, позволяя подклассам изменять тип создаваемых объектов.

Применимость:
1) когда заранее неизвестны типы и зависимости объектов, с которыми должен работать код;
2) когда нужно дать возможность пользователям расширять части нашего фреймворка или библиотеки;
3) когда нужно экономить системные ресурсы, повторно используя уже созданные объекты, вместо создания новых.

Плюсы:
- избавляет объект от привязки к конкретным объектам продуктов;
- выделяет код производства продуктов в одно место, упрощая поддержку кода;
- упрощает добавление новых продуктов в программу;
- реализует принцип открытости/закрытости.

Минусы:
- может привести к созданию больших параллельных иерархий классов, так как для каждого класса продукта надо создать свой
подкласс создателя.

Примеры:
1) Довольно ярким примером может послужить служба доставки товаров. В начале доставка осуществлялась только в границах
одного города посредством курьера на машине, потом появилась необходимость расшириться на всю страну. В данном случае
возникнет проблема, так как весь код уже завязан на объекте машины, а на ней по стране доставка очень долгая и дорогая.
В таком случае, можно выделять общий интерфейс транспорта transport (например, с одним методом deliver()) и реализовывать
его на конкретных экземплярах. Таким образом мы можем выделить конструктор транспорта, который будет возвращать разные
типы (машина, самолет, поезд и т.д.) в зависимости от потребности.

Реализация:
Допустим, у нас магазин по продаже мебели. Сначала мы торговали только шкафами, но потом решили расширить ассортимент
кухнями. Мы могли бы просто добавить большие условные конструкции. Но с каждым
новым товаром читать код становится все сложнее. Помочь может паттерн "Фабричный метод". Мы выделяем интерфейс товара
furniture и реализуем его в каждом объекте. После создаем конструктор, который на вход получает необходимый тип товара
и возвращает новый экземпляр.
*/

type Furniture interface {
	GetDetails() map[string]string
	GetProductId() int
	GetType() string
}

// BaseFurnitureType - базовое поведение наших товаров
type BaseFurnitureType struct {
	productId int
	modelName string
	brand     string
	style     string
}

func (b *BaseFurnitureType) GetDetails() map[string]string {
	res := make(map[string]string, 3)
	res["model"] = b.modelName
	res["brand"] = b.brand
	res["style"] = b.style
	return res
}

func (b *BaseFurnitureType) GetProductId() int {
	return b.productId
}

// реализация конкретных типов товаров

type Closet struct {
	BaseFurnitureType
}

func NewCloset(model, brand, style string, id int) *Closet {
	return &Closet{
		BaseFurnitureType: BaseFurnitureType{
			productId: id,
			modelName: model,
			brand:     brand,
			style:     style,
		},
	}
}

func (c *Closet) GetType() string {
	return "closet"
}

type KitchenSet struct {
	BaseFurnitureType
}

func NewKitchenSet(model, brand, style string, id int) *KitchenSet {
	return &KitchenSet{
		BaseFurnitureType: BaseFurnitureType{
			productId: id,
			modelName: model,
			brand:     brand,
			style:     style,
		},
	}
}

func (k *KitchenSet) GetType() string {
	return "kitchen"
}

// NewFurniture - наша фабрика
func NewFurniture(furnitureType string) Furniture {
	switch furnitureType {
	default:
		return nil
	case "closet":
		return NewCloset("defaultModel", "defaultBrand", "defaultStyle", 0)
	case "kitchen":
		return NewKitchenSet("defaultModel", "defaultBrand", "defaultStyle", 0)
	}
}
