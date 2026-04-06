import { useEffect, useState } from "react";
import About from "../components/About";
import Contact from "../components/Contact";
import Hero from "../components/Hero";
import Loading from "../components/Loading";
import MyPicks from "../components/MyPicks";
import Navigation from "../components/Navigation";
import Work from "../components/Work";

export default function HomePage() {
  const [isLoading, setIsLoading] = useState(true);

  useEffect(() => {
    const timeout = window.setTimeout(() => {
      setIsLoading(false);
    }, 1700);

    return () => window.clearTimeout(timeout);
  }, []);

  return (
    <>
      <Loading isVisible={isLoading} />

      <div className={`experience-shell${isLoading ? " is-loading" : ""}`}>
        <Navigation />

        <div className="experience-main">
          <Hero />
          <Work />
          <MyPicks />
          <About />
          <Contact />
        </div>
      </div>
    </>
  );
}
