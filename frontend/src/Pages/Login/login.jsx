// import * as React from "react";
import React, { useRef, useContext } from "react";
import { useNavigate } from "react-router-dom";
import AuthContext from "../../store/auth-context";

import "./style.css";
import banner from "../../Assets/Images/Login-Images.png";

<<<<<<< HEAD
export const Login = () => {
	const navigate = useNavigate();

	const emailInputRef = useRef();
	const passwordInputRef = useRef();

	const authCtx = useContext(AuthContext);

	const isLoggedIn = authCtx.isLoggedIn;

	if (!isLoggedIn) {
		navigate("/HomePage");
	}

	const handleSubmit = (event) => {
		event.preventDefault();

		const enteredEmail = emailInputRef.current.value;
		const enteredPassword = passwordInputRef.current.value;

		let url;
		url = "https://d440-120-188-38-92.ap.ngrok.io/auth/sign-in";
		fetch(url, {
			method: "POST",
			body: JSON.stringify({
				email: enteredEmail,
				password: enteredPassword,
				returnSecureToken: true,
			}),
			headers: {
				Accept: "/",
				"Content-Type": "application/json",
			},
		})
			.then((res) => {
				if (res.ok) {
					return res.json();
				} else {
					return res.json().then((data) => {
						let errorMassage = "Authentication failed!";
						// if (data && data.error && data.error.message) {
						// 	errorMassage = data.error.message;
						// }

						throw new Error(errorMassage);
					});
				}
			})
			.then((data) => {
				// console.log(data);
				authCtx.login(data.idToken);
			})
			.catch((err) => {
				// alert(err.message);
				alert("Data yang anda masukkan salah, silahkan di cek!");
			});
	};

	return (
		<section id="login-pages">
			<div className="row">
				<div className="login-left col-lg-6">
					<img className="w-75 img-fluid mx-auto" src={banner} alt="banner" />
				</div>
				<div className="login-right col-lg-6 my-auto">
					<p>Welcome Back</p>
					<h2 className="mb-5">Login to your Account</h2>
					<form onSubmit={handleSubmit}>
						<div className="mb-3">
							<label for="emailInput" className="form-label">
								Email
							</label>
							<input
								type="email"
								id="email"
								className="form-control"
								ref={emailInputRef}
							/>
						</div>
						<div className="mb-3">
							<label for="exampleInputPassword1" className="form-label">
								Password
							</label>
							<input
								type="password"
								id="password"
								className="form-control"
								ref={passwordInputRef}
							/>
						</div>
						<button type="submit" className="btn mt-3">
							Login
						</button>
					</form>
				</div>
			</div>
		</section>
	);
=======
const Login = () => {
  const [email, setemail] = useState("");
  const [password, setpassword] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const response = await axios.post(
        "https://d440-120-188-38-92.ap.ngrok.io/auth/sign-in",
        {
          email: email,
          password: password,
        },
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
          },
        }
      );
      console.log(response);
    } catch (error) {
      console.log(error);
    }

    // axios
    // 	.post(
    // 		`https://7dbd-114-10-64-47.ap.ngrok.io/auth/sign-in`,
    // 		{
    // 			email: email,
    // 			password: password,
    // 		},
    // 		{
    // 			headers: {
    // 				Accept: "/",
    // 				"User-Agent": "Thunder Client (https://www.thunderclient.com/)",
    // 				"Content-Type": "application/json",
    // 			},
    // 		}
    // 	)
    // 	.then(function (response) {
    // 		// console.log(username, password);
    // 	});
  };

  return (
    <section id="login-pages">
      <div className="row">
        <div className="login-left col-lg-6">
          <img className="w-75 img-fluid mx-auto" src={banner} alt="banner" />
        </div>
        <div className="login-right col-lg-6 my-auto">
          <p>Welcome Back</p>
          <h2 className="mb-5">Login to your Account</h2>
          <form>
            <div className="mb-3">
              <label for="emailInput" className="form-label">
                Email
              </label>
              <input
                type="email"
                value={email}
                className="form-control"
                onChange={(e) => setemail(e.target.value)}
              />
            </div>
            <div className="mb-3">
              <label for="exampleInputPassword1" className="form-label">
                Password
              </label>
              <input
                type="password"
                value={password}
                className="form-control"
                onChange={(e) => setpassword(e.target.value)}
              />
            </div>
            {/* <div className="mb-3 form-check">
							<input
								type="checkbox"
								className="form-check-input"
								id="exampleCheck1"
							/>
							<label className="form-check-label" for="exampleCheck1">
								Check me out
							</label>
						</div> */}
            {/* <Link to="/home"> */}
            <button type="submit" className="btn mt-3" onClick={handleSubmit}>
              Login
            </button>
            {/* </Link> */}
          </form>
        </div>
      </div>
    </section>
  );
>>>>>>> e773c085e77b42a1461e27b8a416f0bf9c59e89b
};

export default Login;
