import React from "react";
import { Link } from "react-router-dom";

import Logo from "../../../Assets/Images/Logo.png";

import "./style.css";
import { Container, Navbar, Nav, Row, Col } from "react-bootstrap";

const navbar = () => {
	return (
		<section id="navbar-landing">
			<Navbar className="nav" expand="lg" variant="light">
				<Container>
					<Row>
						<Col>
							<img src={Logo} className="img-fluid" alt="logo" />
						</Col>
					</Row>
					<Row className="right-content">
						<Navbar.Toggle aria-controls="basic-navbar-nav" />
						<Navbar.Collapse id="basic-navbar-nav">
							<Nav className="me-auto">
								<Nav>
									<Link to="/register">
										<button className="signup float-lg-right">Sign Up</button>
									</Link>
								</Nav>
								<Nav>
									<Link to="/login">
										<button className="login float-lg-right mr-2">
											Log In
										</button>
									</Link>
								</Nav>
							</Nav>
						</Navbar.Collapse>
					</Row>
				</Container>
			</Navbar>
		</section>
	);
};

export default navbar;
