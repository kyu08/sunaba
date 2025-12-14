package controllers

import javax.inject.Inject
import javax.inject.Singleton
import play.api.mvc.AbstractController
import play.api.mvc.Action
import play.api.mvc.AnyContent
import play.api.mvc.ControllerComponents
import play.api.mvc.Request

@Singleton
class HelloController @Inject()(cc: ControllerComponents) extends AbstractController(cc) {

  def get(a: Option[String], b: Option[String]) =
    Action { implicit request: Request[AnyContent] =>
      Ok {
        val result = (a, b) match {
          case (Some(x), Some(y)) => Some(x.toInt + y.toInt)
          case (None, Some(_)) => None
          case (Some(_), None) => None
          case (None, None) => None
        }
        result
          .map(r => s"a + ab = $r")
          .getOrElse("""root""")

      }
    }

  def plus(a: Option[String], b: Option[String]) =
    Action { implicit request: Request[AnyContent] =>
      Ok {
        val result = (a, b) match {
          case (Some(x), Some(y)) => Some(x.toInt + y.toInt)
          case (None, Some(_)) => None
          case (Some(_), None) => None
          case (None, None) => None
        }
        result
          .map(r => s"a + bb = $r")
          .getOrElse("""Please give arguments of a and b.""")

      }
    }
}
