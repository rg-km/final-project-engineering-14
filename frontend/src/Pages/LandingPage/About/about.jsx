import React from "react";
import { Link } from "react-router-dom";
import "./style.css";
import Banner from "../../../Assets/Images/Landing-About.png";
import Footer from "../../../Components/Footer/Footer";

const about = () => {
  return (
    <section id="about-pages">
      <div className="container">
        <div className="row">
          <div className="about-left col-lg-6 my-auto">
            <h1 className="mb-3">
              You can join with RuangArah and upgrade your skill for your{" "}
              <span>Bright Future.</span>
            </h1>
            <p className="mb-5">
              Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pharetra
              eget sit vel egestas volutpat sed. Turpis ut egestas ridiculus a
              egestas ultrices. Quis proin tortor, elementum pretium interdum.
            </p>
            <Link to="/login">
              <button>Start Learning</button>
            </Link>
          </div>
          <div className="about-right col-lg-6">
            <img
              className="img-fluid mx-auto my-auto"
              src={Banner}
              alt="banner"
            />
          </div>
        </div>
      </div>
      <Footer />
    </section>
  );
};

export default about;
