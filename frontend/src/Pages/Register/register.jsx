import React from "react";
import "./register.css";
import banner from "../../Assets/Images/Register-Images.png";
import { useState } from "react";
import axios from "axios";
// import { Link } from "react-router-dom";

export default function Register() {
  const [username, setUsername] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [phone, setPhone] = useState("");

<<<<<<< HEAD
	const handleSubmit = async (e) => {
		e.preventDefault();
		try {
			let res = await axios.post(
				"https://d440-120-188-38-92.ap.ngrok.io/auth/sign-up",
				{
					username: username,
					email: email,
					password: password,
					phone: phone,
				},
				{
					headers: {
						Accept: "/",
						"Content-Type": "application/json",
					},
				}
			);
			console.log(res);
		} catch (error) {
			console.log(error);
		}
	};
	return (
		<section id="register-pages">
			<div className="row">
				<div className="register-left col-lg-6">
					<img className="w-75 img-fluid mx-auto" src={banner} alt="banner" />
				</div>
				<div className="register-right col-lg-6 my-auto">
					<h2 className="mb-5">Register Account</h2>
					<h4>Make your precious time more efficient</h4>
					<p>
						Let's get you all set up so you can verify your personal account{" "}
						<br /> and begin setting up your profile.
					</p>
					<form onSubmit={handleSubmit} className="mt-4">
						<div className="form-group">
							<label htmlFor="" className="form-label">
								Username
							</label>
							<input
								type="text"
								className="form-control"
								id="username"
								name="username"
								onChange={(e) => setUsername(e.target.value)}
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
								onChange={(e) => setEmail(e.target.value)}
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
								onChange={(e) => setPassword(e.target.value)}
							/>
						</div>
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
								onChange={(e) => setPhone(e.target.value)}
							/>
						</div>
						<div className="d-grid gap-2 mt-3">
							{/* <Link to="/login"> */}
							<button
								type="submit"
								className="btn-color-theme btn-lg btn-block p-2 mt-3"
							>
								Register
							</button>
							{/* </Link> */}
						</div>
						<p className="have-account mt-4">
							Already have an account?{" "}
							<a href="/login" className="text-decoration-none">
								Login
							</a>
						</p>
					</form>
				</div>
			</div>
		</section>
	);
=======
  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let res = await axios.post(
        "https://d440-120-188-38-92.ap.ngrok.io/auth/sign-up",
        {
          username: username,
          email: email,
          password: password,
          phone: phone,
        },
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
          },
        }
      );
      console.log(res);
    } catch (error) {
      console.log(error);
    }
  };
  return (
    <section id="register-pages">
      <div className="row">
        <div className="register-left col-lg-6">
          <img className="w-75 img-fluid mx-auto" src={banner} alt="banner" />
        </div>
        <div className="register-right col-lg-6 my-auto">
          <h2 className="mb-5">Register Account</h2>
          <h4>Make your precious time more efficient</h4>
          <p>
            Let's get you all set up so you can verify your personal account{" "}
            <br /> and begin setting up your profile.
          </p>
          <form onSubmit={handleSubmit} className="mt-4">
            <div className="form-group">
              <label htmlFor="" className="form-label">
                Username
              </label>
              <input
                type="text"
                className="form-control"
                id="username"
                name="username"
                onChange={(e) => setUsername(e.target.value)}
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
                onChange={(e) => setEmail(e.target.value)}
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
                onChange={(e) => setPassword(e.target.value)}
              />
            </div>
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
                onChange={(e) => setPhone(e.target.value)}
              />
            </div>
            <div className="d-grid gap-2 mt-3">
              {/* <Link to="/login"> */}
              <button
                type="submit"
                className="btn-color-theme btn-lg btn-block p-2 mt-3"
              >
                Register
              </button>
              {/* </Link> */}
            </div>
            <p className="have-account mt-4">
              Already have an account?{" "}
              <a href="/login" className="text-decoration-none">
                Login
              </a>
            </p>
          </form>
        </div>
      </div>
    </section>
  );
>>>>>>> e773c085e77b42a1461e27b8a416f0bf9c59e89b
}
