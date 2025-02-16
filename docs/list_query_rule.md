# 开列表查询规则

## 通用列表查询请求

| 字段名       | 类型        | 格式                                  | 字段描述    | 示例                                                                                                       | 备注                                                               |
|-----------|-----------|-------------------------------------|---------|----------------------------------------------------------------------------------------------------------|------------------------------------------------------------------|
| page      | `number`  |                                     | 当前页码    |                                                                                                          | 默认为`1`，最小值为`1`。                                                  |
| pageSize  | `number`  |                                     | 每页的行数   |                                                                                                          | 默认为`10`，最小值为`1`。                                                 |
| query     | `string`  | `json object` 或 `json object array` | AND过滤条件 | json字符串: `{"field1":"val1","field2":"val2"}` 或者`[{"field1":"val1"},{"field1":"val2"},{"field2":"val2"}]` | `map`和`array`都支持，当需要同字段名，不同值的情况下，请使用`array`。具体规则请见：[过滤规则](#过滤规则) |
| or        | `string`  | `json object` 或 `json object array` | OR过滤条件  | 同 AND过滤条件                                                                                                |                                                                  |
| orderBy   | `string`  | `json string array`                 | 排序条件    | json字符串：`["-create_time", "type"]`                                                                       | json的`string array`，字段名前加`-`是为降序，不加为升序。具体规则请见：[排序规则](#排序规则)      |
| noPaging  | `boolean` |                                     | 是否不分页   |                                                                                                          | 此字段为`true`时，`page`、`pageSize`字段的传入将无效用。                          |
| fieldMask | `string`  | 其语法为使用逗号分隔字段名                       | 字段掩码    | 例如：id,realName,userName。                                                                                 | 此字段是`SELECT`条件，为空的时候是为`*`。                                       |

## 排序规则

排序操作本质上是`SQL`里面的`Order By`条件。

| 序列 | 示例                 | 备注           |
|----|--------------------|--------------|
| 升序 | `["type"]`         |              |
| 降序 | `["-create_time"]` | 字段名前加`-`是为降序 |

## 过滤规则

过滤器操作本质上是`SQL`里面的`WHERE`条件。

过滤器的规则，遵循了Python的ORM的规则，比如：

- [Tortoise ORM Filtering](https://tortoise.github.io/query.html#filtering)。
- [Django Field lookups](https://docs.djangoproject.com/en/4.2/ref/models/querysets/#field-lookups)

如果只是普通的查询，只需要传递`字段名`即可，但是如果需要一些特殊的查询，那么就需要加入`操作符`了。

特殊查询的语法规则其实很简单，就是使用双下划线`__`分割字段名和操作符：

```text
{字段名}__{查找类型} : {值}
{字段名}.{JSON字段名}__{查找类型} : {值}
```

| 查找类型        | 示例                                                            | SQL                                                                                                                                                                                                                       | 备注                                                                                                            |
|-------------|---------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|---------------------------------------------------------------------------------------------------------------|
| not         | `{"name__not" : "tom"}`                                       | `WHERE NOT ("name" = "tom")`                                                                                                                                                                                              |                                                                                                               |
| in          | `{"name__in" : "[\"tom\", \"jimmy\"]"}`                       | `WHERE name IN ("tom", "jimmy")`                                                                                                                                                                                          |                                                                                                               |
| not_in      | `{"name__not_in" : "[\"tom\", \"jimmy\"]"}`                   | `WHERE name NOT IN ("tom", "jimmy")`                                                                                                                                                                                      |                                                                                                               |
| gte         | `{"create_time__gte" : "2023-10-25"}`                         | `WHERE "create_time" >= "2023-10-25"`                                                                                                                                                                                     |                                                                                                               |
| gt          | `{"create_time__gt" : "2023-10-25"}`                          | `WHERE "create_time" > "2023-10-25"`                                                                                                                                                                                      |                                                                                                               |
| lte         | `{"create_time__lte" : "2023-10-25"}`                         | `WHERE "create_time" <= "2023-10-25"`                                                                                                                                                                                     |                                                                                                               |
| lt          | `{"create_time__lt" : "2023-10-25"}`                          | `WHERE "create_time" < "2023-10-25"`                                                                                                                                                                                      |                                                                                                               |
| range       | `{"create_time__range" : "[\"2023-10-25\", \"2024-10-25\"]"}` | `WHERE "create_time" BETWEEN "2023-10-25" AND "2024-10-25"` <br>或<br> `WHERE "create_time" >= "2023-10-25" AND "create_time" <= "2024-10-25"`                                                                             | 需要注意的是: <br>1. 有些数据库的BETWEEN实现的开闭区间可能不一样。<br>2. 日期`2005-01-01`会被隐式转换为：`2005-01-01 00:00:00`，两个日期一致就会导致查询不到数据。 |
| isnull      | `{"name__isnull" : "True"}`                                   | `WHERE name IS NULL`                                                                                                                                                                                                      |                                                                                                               |
| not_isnull  | `{"name__not_isnull" : "False"}`                              | `WHERE name IS NOT NULL`                                                                                                                                                                                                  |                                                                                                               |
| contains    | `{"name__contains" : "L"}`                                    | `WHERE name LIKE '%L%';`                                                                                                                                                                                                  |                                                                                                               |
| icontains   | `{"name__icontains" : "L"}`                                   | `WHERE name ILIKE '%L%';`                                                                                                                                                                                                 |                                                                                                               |
| startswith  | `{"name__startswith" : "La"}`                                 | `WHERE name LIKE 'La%';`                                                                                                                                                                                                  |                                                                                                               |
| istartswith | `{"name__istartswith" : "La"}`                                | `WHERE name ILIKE 'La%';`                                                                                                                                                                                                 |                                                                                                               |
| endswith    | `{"name__endswith" : "a"}`                                    | `WHERE name LIKE '%a';`                                                                                                                                                                                                   |                                                                                                               |
| iendswith   | `{"name__iendswith" : "a"}`                                   | `WHERE name ILIKE '%a';`                                                                                                                                                                                                  |                                                                                                               |
| exact       | `{"name__exact" : "a"}`                                       | `WHERE name LIKE 'a';`                                                                                                                                                                                                    |                                                                                                               |
| iexact      | `{"name__iexact" : "a"}`                                      | `WHERE name ILIKE 'a';`                                                                                                                                                                                                   |                                                                                                               |
| regex       | `{"title__regex" : "^(An?\|The) +"}`                          | MySQL: `WHERE title REGEXP BINARY '^(An?\|The) +'`  <br> Oracle: `WHERE REGEXP_LIKE(title, '^(An?\|The) +', 'c');`  <br> PostgreSQL: `WHERE title ~ '^(An?\|The) +';`  <br> SQLite: `WHERE title REGEXP '^(An?\|The) +';` |                                                                                                               |
| iregex      | `{"title__iregex" : "^(an?\|the) +"}`                         | MySQL: `WHERE title REGEXP '^(an?\|the) +'`  <br> Oracle: `WHERE REGEXP_LIKE(title, '^(an?\|the) +', 'i');`  <br> PostgreSQL: `WHERE title ~* '^(an?\|the) +';`  <br> SQLite: `WHERE title REGEXP '(?i)^(an?\|the) +';`   |                                                                                                               |
| search      |                                                               |                                                                                                                                                                                                                           |                                                                                                               |

以及将日期提取出来的查找类型：

| 查找类型         | 示例                                   | SQL                                               | 备注                   |
|--------------|--------------------------------------|---------------------------------------------------|----------------------|
| date         | `{"pub_date__date" : "2023-01-01"}`  | `WHERE DATE(pub_date) = '2023-01-01'`             |                      |
| year         | `{"pub_date__year" : "2023"}`        | `WHERE EXTRACT('YEAR' FROM pub_date) = '2023'`    | 哪一年                  |
| iso_year     | `{"pub_date__iso_year" : "2023"}`    | `WHERE EXTRACT('ISOYEAR' FROM pub_date) = '2023'` | ISO 8601 一年中的周数      |
| month        | `{"pub_date__month" : "12"}`         | `WHERE EXTRACT('MONTH' FROM pub_date) = '12'`     | 月份，1-12              |
| day          | `{"pub_date__day" : "3"}`            | `WHERE EXTRACT('DAY' FROM pub_date) = '3'`        | 该月的某天(1-31)          |
| week         | `{"pub_date__week" : "7"}`           | `WHERE EXTRACT('WEEK' FROM pub_date) = '7'`       | ISO 8601 周编号 一年中的周数 |
| week_day     | `{"pub_date__week_day" : "tom"}`     | ``                                                | 星期几                  |
| iso_week_day | `{"pub_date__iso_week_day" : "tom"}` | ``                                                |                      |
| quarter      | `{"pub_date__quarter" : "1"}`        | `WHERE EXTRACT('QUARTER' FROM pub_date) = '1'`    | 一年中的季度              |
| time         | `{"pub_date__time" : "12:59:59"}`    | ``                                                |                      |
| hour         | `{"pub_date__hour" : "12"}`          | `WHERE EXTRACT('HOUR' FROM pub_date) = '12'`      | 小时(0-23)             |
| minute       | `{"pub_date__minute" : "59"}`        | `WHERE EXTRACT('MINUTE' FROM pub_date) = '59'`    | 分钟 (0-59)            |
| second       | `{"pub_date__second" : "59"}`        | `WHERE EXTRACT('SECOND' FROM pub_date) = '59'`    | 秒 (0-59)             |

## 项目代码

* [kratos-admin Gitee](https://gitee.com/tx7do/kratos-admin)
* [kratos-admin Github](https://github.com/tx7do/kratos-admin)
