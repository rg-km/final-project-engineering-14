import React from "react";
import "./register.css";
import banner from "../../Assets/Images/Register-Images.png";
import { BsFillEyeFill } from "react-icons/bs";

export default function register() {
  return (
    <section id="register-pages">
      <div className="row">
        <div className="register-left col-md-4">
          <img className="register-logo img-fluid" src={banner} alt="banner" />
        </div>
        <div className="register-right col-md-5">
          <div className="card shadow-sm p-3 border-radius-10">
            <div className="card-body">
              <h2>Register Account</h2>
              <div className="register-text mt-5">
                <h3>Make your precious time more efficient</h3>
                <p>
                  Let's get you all set up so you can verify your personal
                  account and begin setting up your profile.
                </p>
                <form action="" className="mt-4">
                  <div className="form-group">
                    <label htmlFor="" className="form-label">
                      Username
                    </label>
                    <input
                      type="text"
                      className="form-control"
                      id="name"
                      name="name"
                      placeholder="Enter your username ..."
                    />
                  </div>
                  <div className="form-group mt-3">
                    <label htmlFor="" className="form-label">
                      Email
                    </label>
                    <input
                      type="text"
                      className="form-control"
                      id="email"
                      name="email"
                      placeholder="Enter your email ..."
                    />
                  </div>
                  <div className="form-group mt-3">
                    {/* input-group*/}
                    <label htmlFor="" className="form-label">
                      Password
                    </label>
                    <input
                      type="password"
                      className="form-control"
                      id="password"
                      name="password"
                      placeholder="Enter your password ..."
                    />
                  </div>
                  {/* <div className="input-group-opened">
                      <span className="input-group-text">
                        <i>
                          <BsFillEyeFill />
                        </i>
                      </span>
                    </div> */}
                  <div className="form-group mt-3">
                    {/* input-group*/}
                    <label htmlFor="" className="form-label">
                      Phone Number
                    </label>
                    <input
                      type="text"
                      className="form-control"
                      id="PhoneNumber"
                      name="PhoneNumber"
                      placeholder="Enter your Phone Number ..."
                    />
                  </div>
                  <div className="d-grid gap-2 mt-3">
                    <button
                      type="submit"
                      className="btn-color-theme btn-lg btn-block p-2 mt-3"
                    >
                      Register
                    </button>
                  </div>
                  <p className="text-center mt-2">
                    Already have an account?{" "}
                    <a href="/login" className="text-decoration-none">
                      Login
                    </a>
                  </p>
                </form>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  );
}
