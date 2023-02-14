package calculator;

import java.time.Duration;
import java.util.*;

import io.gatling.javaapi.core.*;
import io.gatling.javaapi.http.*;
import io.gatling.javaapi.jdbc.*;

import static io.gatling.javaapi.core.CoreDsl.*;
import static io.gatling.javaapi.http.HttpDsl.*;
import static io.gatling.javaapi.jdbc.JdbcDsl.*;

public class CalculatorStressTestSimulation extends Simulation {

    {
        HttpProtocolBuilder httpProtocol = http
          .baseUrl("http://localhost:8000/calc")
          .inferHtmlResources(AllowList(), DenyList(".*\\.js", ".*\\.css", ".*\\.gif", ".*\\.jpeg", ".*\\.jpg", ".*\\.ico", ".*\\.woff", ".*\\.woff2", ".*\\.(t|o)tf", ".*\\.png", ".*detectportal\\.firefox\\.com.*"))
          .acceptHeader("text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
          .acceptEncodingHeader("gzip, deflate")
          .acceptLanguageHeader("en-US,en;q=0.9,pt;q=0.8")
          .upgradeInsecureRequestsHeader("1")
          .userAgentHeader("Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36");
        
        Map<CharSequence, String> headers_0 = new HashMap<>();
        headers_0.put("Proxy-Connection", "keep-alive");
        
        Map<CharSequence, String> headers_11 = new HashMap<>();
        headers_11.put("Cache-Control", "max-age=0");
        headers_11.put("Origin", "http://localhost:8000/calc");
        headers_11.put("Proxy-Connection", "keep-alive");
    
    
        ScenarioBuilder scn = scenario("CalculatorSimulation")
          .exec(
            http("request_0")
              .get("/history")
              .headers(headers_0)
          )
          .pause(3)
          .exec(
            http("request_1")
              .get("/div/133/6986")
              .headers(headers_0)
          )
          .pause(3)
          .exec(
            http("request_2")
              .get("/history")
              .headers(headers_0)
          );
    
          setUp(scn.injectOpen(atOnceUsers(2000), rampUsers(2500).during(15))).protocols(httpProtocol);
      }
}