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

	const handleSubmit = async (e) => {
		e.preventDefault();
		try {
			let res = await axios.post(
				"https://f61a-120-188-92-248.ap.ngrok.io/auth/sign-up",
				{
					username: username,
					email: email,
					password: password,
					phone: phone,
				},
				{
					headers: {
						Accept: "/",
						"User-Agent": "Thunder Client (https://www.thunderclient.com/)",
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
				<div className="register-left col-md-4">
					<img className="register-logo img-fluid" src={banner} alt="banner" />
				</div>
				<div className="register-right col-md-5">
					<div className="card  p-3 border-radius-10">
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
											placeholder="Enter your email ..."
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
											placeholder="Enter your password ..."
											onChange={(e) => setPassword(e.target.value)}
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
											onChange={(e) => setPhone(e.target.value)}
										/>
									</div>
									<div className="d-grid gap-2 mt-3">
										{/* <Link to="/login"> */}
										<button
											type="submit"
											className="btn-color-theme btn-lg btn-block p-2 mt-3"
											onClick={handleSubmit}
										>
											Register
										</button>
										{/* </Link> */}
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
