package controllers


func (this *MainController) ApiV1Get() {

	this.Data["json"] = map[string]interface{}{
		"Plugin":   "Testplugin",
		"Level":    "3",
		"Duration": "5",
	}
	this.ServeJSON()

}
