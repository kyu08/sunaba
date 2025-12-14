package controllers

import javax.inject.Inject
import play.api.libs.json.{JsValue, Json}
import play.api.mvc._

class HelloController @Inject() (cc: ControllerComponents) extends AbstractController(cc) {

  def hello(): Action[AnyContent] = {
    val result: Result = Ok("Hello World!")
    Action(result)
  }

  def helloJson(): Action[AnyContent] = Action {
    val json: JsValue =
      Json.obj("hello" -> "world", "language" -> "scala")
    Ok(json)
  }

}