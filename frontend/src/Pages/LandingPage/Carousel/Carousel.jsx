import React from "react";
import { Link } from "react-router-dom";
import "./style.css";
import Background from "../../../Assets/Images/Landing-Background-2.png";

const Carousel = () => {
	return (
		<section id="carousel-landing">
			<div className="carousel-inner">
				<div className="carousel-item active">
					<img src={Background} className="d-block w-100 gmbr" alt="..." />
					<div className="carousel-caption d-none d-md-block">
						<h2>
							Build Your Skills With{" "}
							<span>
								{" "}
								Experts <br /> Any Time Anywhere
							</span>{" "}
						</h2>
						<p className="mt-3 mb-5">
							Free online courses from the world’s Leading experts. join 10+
							million Learners today.
						</p>
						<Link to="/register">
							<button>Join Us Today</button>
						</Link>
					</div>
				</div>
			</div>
		</section>
	);
};

export default Carousel;
