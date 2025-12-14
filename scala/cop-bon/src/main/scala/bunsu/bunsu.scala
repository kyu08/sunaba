package bunsu

object bunsu {
  def main(args: Array[String]): Unit = {
    implicit def intToRational(i: Int) = new Rational(i)
    println("-----------------------")
    println(1 + new Rational(1,24))
    println(new Rational(1, 2) - new Rational(3, 2))
    println("-----------------------")
  }
}
