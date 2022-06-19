import React from "react";
import { Link } from "react-router-dom";
import "./style.css";
import Background from "../../../Assets/Images/Home-Background-2.png";

function Main() {
  return (
    <section id="main-pages">
      <div className="row">
        <div className="main-left col-md-6">
          <img className="main-logo img-fluid" src={Background} alt="banner" />
        </div>
        <div className="main-right col-md-5">
          <div className="card p-3 border-radius-10">
            <div className="main-text card-body">
              <h1>Start Your Journey</h1>
              <div className="mt-5">
                <p>Do You Know What You Must Learn Next ?</p>
              </div>
            </div>
            <div className="button">
              <Link to="/#">
                <button
                  type="submit"
                  className="btn-color-theme btn-lg  p-2 mt-3 mb-5"
                  //   onClick={handleSubmit}
                >
                  Start Here
                </button>
              </Link>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}

export default Main;
