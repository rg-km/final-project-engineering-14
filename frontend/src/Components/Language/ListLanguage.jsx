import React from "react";
import "./Language.css";
import { Link } from "react-router-dom";

//React BS
import Card from "react-bootstrap/Card";

export default function ListLanguage({ image, language }) {
  return (
    <section id="list-language">
      <Card>
        <Link to="/Question" className="pages-link">
          <Card.Img variant="top" src={image} className="image-language" />
          <Card.Body>
            <Card.Text>
              <div className="card-text text-center">
                <h2>{language}</h2>
              </div>
            </Card.Text>
          </Card.Body>
        </Link>
      </Card>
    </section>
  );
}
