import React from "react";
import "./language.css";
import Bahasa1 from "../../Assets/Images/Bahasa1.png";
import Bahasa2 from "../../Assets/Images/Bahasa2.png";
import Bahasa3 from "../../Assets/Images/Bahasa3.png";
import Bahasa4 from "../../Assets/Images/Bahasa4.png";
import Bahasa5 from "../../Assets/Images/Bahasa5.png";
import Bahasa6 from "../../Assets/Images/Bahasa6.png";
import Bahasa7 from "../../Assets/Images/Bahasa7.png";
import Bahasa8 from "../../Assets/Images/Bahasa8.png";

const language = () => {
	return (
		<section id="language-page">
			<div className="container">
				<h1>
					Bahasa <span> Pemrograman </span>{" "}
				</h1>
				<div className="card">
					<div className="row mx-auto">
						<div className="col-4">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa1} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>Go</h2>
								</div>
							</div>
						</div>
						<div className="col-4">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa2} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>Python</h2>
								</div>
							</div>
						</div>
						<div className="col-4">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa3} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>Ruby</h2>
								</div>
							</div>
						</div>
					</div>
					<div className="row mx-auto">
						<div className="col-4">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa4} className="img-fluid " alt="" />
								</div>
								<div className="card-text text-center">
									<h2>PHP</h2>
								</div>
							</div>
						</div>
						<div className="col-4">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa5} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>Rust</h2>
								</div>
							</div>
						</div>
						<div className="col-4">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa6} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>C#</h2>
								</div>
							</div>
						</div>
					</div>
					<div className="row mx-auto">
						<div className="col-6">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa7} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>Java</h2>
								</div>
							</div>
						</div>
						<div className="col-6">
							<div className="card mt-5">
								<div className="card-img">
									<img src={Bahasa8} className="img-fluid" alt="" />
								</div>
								<div className="card-text text-center">
									<h2>Javascript</h2>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</section>
	);
};

export default language;
