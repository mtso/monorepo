namespace middleware_dotnet
{
    public interface IDemoService {
        int GetNum();
    }

    public class DemoService: IDemoService {
        public int GetNum() {
            return 42;
        }
    }
}
