import React from "react";
import "./Language.css";
import { Link } from "react-router-dom";

//React BS
import Card from "react-bootstrap/Card";

export default function ListLanguage({ id, image, language }) {
	return (
		<section id="list-language">
			<Card>
				<Link to="/Question" className="pages-link">
					<Card.Img variant="top" src={image} className="image-language" />
					<Card.Body>
						<Card.Text className="card-text text-center">{language}</Card.Text>
					</Card.Body>
				</Link>
			</Card>
		</section>
	);
}
