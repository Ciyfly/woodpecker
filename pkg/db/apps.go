/*
 * @Date: 2022-04-18 10:14:46
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:14:53
 */
/*
 * @Date: 2022-04-18 10:12:24
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:13:38
 */
/*
 * @Date: 2022-04-18 10:05:08
 * @LastEditors: recar
 * @LastEditTime: 2022-04-18 10:12:02
 */

package db

type App struct {
	Id       uint `gorm:"primaryKey"`
	AppName  string
	Provider string
	Desc     string
}
