/**
 * @Author: lirn
 * @Description:
 * @File: validation_service_provider
 * @Date: 2022/12/29 14:41
 */
package providers

import (
	"github.com/goravel/framework/contracts/validation"
	"github.com/goravel/framework/facades"
)

type ValidationServiceProvider struct {
}

func (receiver *ValidationServiceProvider) Register() {

}

func (receiver *ValidationServiceProvider) Boot() {
	if err := facades.Validation.AddRules(receiver.rules()); err != nil {
		facades.Log.Errorf("add rules error: %+v", err)
	}
}

func (receiver *ValidationServiceProvider) rules() []validation.Rule {
	return []validation.Rule{}
}
