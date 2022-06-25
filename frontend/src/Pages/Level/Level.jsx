import React from "react";
import { Link } from "react-router-dom";
import Footer from "../../Components/Footer/Footer";
import "./Level.css";
import Banner from "../../Assets/Images/Level-image.png";

//React Bs
import Card from "react-bootstrap/Card";
import Image from "react-bootstrap/Image";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";

export default function Level() {
	return (
		<section id="level-pages">
			<Container>
				<Row>
					<Col sm={3}>
						<div className="level-image pt-5">
							<Image src={Banner} fluid />
						</div>
					</Col>
					<Col sm={8}>
						<div className="level-body pt-2">
							<h1>
								Congratulation! <span> Your Level Is</span>{" "}
							</h1>
						</div>
						<div className="card-rank">
							<Card>
								<Card.Body>
									<Card.Text>INTERMEDIATE</Card.Text>
								</Card.Body>
							</Card>
						</div>
						<div className="recommendation">
							<Link to="/Recomendation">
								<button>SEE RECOMMENDATION</button>
							</Link>
						</div>
					</Col>
				</Row>
			</Container>
		</section>
	);
}
