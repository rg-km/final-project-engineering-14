import React from "react";
import { Link } from "react-router-dom";
import "./style.css";
import Background from "../../../Assets/Images/Home-Background-2.png";

import { Col, Row } from "react-bootstrap";

function Main() {
	return (
		<section id="main-pages">
			<Row>
				<Col className="main-left col-md-6">
					<img
						className="main-logo img-sm-fluid"
						src={Background}
						alt="banner"
					/>
				</Col>
				<Col className="main-right col-md-5">
					<div className="card border-radius-10">
						<div className="main-text card-body">
							<h1>Start Your Journey</h1>
							<div className="mt-5">
								<p>Do You Know What You Must Learn Next ?</p>
							</div>
						</div>
						<div className="button">
							<Link to="/Language">
								<button
									type="submit"
									className="btn-color-theme btn-lg  p-2 mt-3 mb-5"
								>
									Start Here
								</button>
							</Link>
						</div>
					</div>
				</Col>
			</Row>
		</section>
	);
}

export default Main;
