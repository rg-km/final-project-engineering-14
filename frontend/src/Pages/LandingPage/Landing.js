import React from "react";
import Navbar from "./Navbar/navbar";
import Carousel from "./Carousel/Carousel";
import About from "./About/about";
import OurTeams from "./OurTeams/OurTeams";

const Landing = () => {
  return (
    <>
      <Navbar />
      <Carousel />
      <About />
      <OurTeams />
    </>
  );
};

export default Landing;
