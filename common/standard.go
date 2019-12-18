/*
@Time : 2019/12/12 11:43 上午
@Author : 伏火
@File : standard.go
@Software: GoLand
*/
package common

/*
	HTTP 状态码
 */
const HTTP_OK = 200
const HTTP_CLIENT_ERROR =201
const HTTP_AUTH_ERROR = 401
const HTTP_NOT_FOUND_ERROR = 404
const HTTP_SERVER_ERROR = 500

/*
	HTTP 返回信息
 */
const HTTP_OK_MESSAGE = "操作成功！"
const HTTP_CLIENT_ERROR_MESSAGE = "参数错误！"
const HTTP_AUTH_ERROR_MESSAGE = "没有权限！"
const HTTP_NOT_FOUND_ERROR_MESSAGE = "无法找到该页面！"
const HTTP_SERVER_ERROR_MESSAGE = "服务器内部错误"