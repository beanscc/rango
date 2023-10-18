package internal

// CountryCode ISO country codes
// 数据来源：ISO 3166-1：https://www.iso.org/obp/ui/#search
// 字段含义说明：https://www.iso.org/glossary-for-iso-3166.html
// https://www.iso.org/iso-3166-country-codes.html
/*
将下面 js 代码放在浏览器 console 模块中执行，获取当前分页的数据
```js
// js 数据解析脚本
let rows = document.querySelectorAll('tbody .v-grid-row');

let data = new Array();
for (let i = 0; i < rows.length; i++) {
    let row = rows.item(i);
    let fields = row.querySelectorAll('td');

    data.push({
        'en_name': fields.item(0).innerText,
        'alpha2_code': fields.item(2).innerText,
        'alpha3_code': fields.item(3).innerText,
        'numeric': fields.item(4).innerText,
    })
}

console.log('country code:', JSON.stringify(data));
```
*/
type CountryCode struct {
	Name    string // english short name
	Alpha2  string // a two-letter code that represents a country name, recommended as the general purpose code
	Alpha3  string // a three-letter code that represents a country name, which is usually more closely related to the country name
	Numeric string // a three-digit numeric code (numeric-3) which can be useful if you need to avoid using Latin script
}
