/*
 * @Author: jianghao
 * @Date: 2022-07-29 09:48:24
 * @LastEditors: jianghao
 * @LastEditTime: 2022-10-17 11:27:44
 */
package init_load

import (
	"reflect"
	"testing"

	"github.com/gofiber/fiber/v2"
)

func TestRouters(t *testing.T) {
	tests := []struct {
		name string
		want *fiber.App
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Routers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Routers() = %v, want %v", got, tt.want)
			}
		})
	}
}
