import React from "react";
import "./register.css";
import banner from "../../Assets/Images/Register-Images.png";

export default function register() {
  return (
    <section id="register-pages">
      <div className="row">
        <div className="register-left col-lg-5 ">
          <img className="register-logo" src={banner} alt="banner" />
        </div>
        <div className="register-right col-lg-6 mx-4">
          <form>
            <h2 className="header-register mb-5">Register Account</h2>

            <div className="register-text">
              <h2>Make your precious time more efficient</h2>
              <p>
                Let's get you all set up so you can verify your personal account
                <br></br>and begin setting up your profile.
              </p>
              <br></br>
            </div>
            <div className="mb-3">
              <label htmlFor="name" className="form-label">
                Username
              </label>
              <input
                type="text"
                className="form-control"
                id="name"
                name="name"
                placeholder="Enter your username ..."
                // onChange={}
              />
            </div>

            <div className="mb-3">
              <label htmlFor="email" className="form-label">
                Email
              </label>

              <input
                type="email"
                className="form-control"
                id="email"
                name="email"
                placeholder="Enter your email ..."
                // onChange={}
              />
            </div>

            <div className="mb-3">
              <label htmlFor="password" className="form-label">
                Password
              </label>
              <input
                type="password"
                className="form-control"
                id="password"
                name="password"
                placeholder="Enter your password ..."
                // onChange={}
              />
            </div>

            <div className="mb-3">
              <label htmlFor="password" className="form-label">
                Phone Number
              </label>
              <input
                type="text"
                className="form-control"
                id="number"
                name="number"
                placeholder="Enter your phone number ..."
                // onChange={}
              />
            </div>

            <button type="submit" className="btn mt-3">
              Register
            </button>
            <div className="login mt-3">
              <p>
                Already have an account? <a href="/login">Login</a>
              </p>
            </div>
          </form>
        </div>
      </div>
    </section>
  );
}
