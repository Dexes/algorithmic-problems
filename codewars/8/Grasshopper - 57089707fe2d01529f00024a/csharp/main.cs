public class Player {
  private int health = 100;

  public int Health
  {
    get { return this.health; }
    set { health = value; }
  }

  public Player() {}

  public bool CheckAlive() { return Health > 0; }
}