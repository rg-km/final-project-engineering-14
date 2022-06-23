import React from "react";
import "./createQuestion.css";

// React Boostraps
import { Container, Form, Button } from "react-bootstrap";

const createQuestion = () => {
	return (
		<section id="create-question">
			<Container className="container-outside bg-dark">
				<Form>
					<Form.Group className="mb-3" controlId="formBasicEmail">
						<Form.Label>Programming Language</Form.Label>
						<Form.Control type="text" />
						<Form.Text className="text-muted">
							We'll never share your email with anyone else.
						</Form.Text>
					</Form.Group>

					<Form.Group className="mb-3" controlId="formBasicPassword">
						<Form.Label>Question</Form.Label>
						<Form.Control as="textarea" />
					</Form.Group>
					<Form.Group className="mb-3" controlId="formBasicCheckbox">
						<Form.Check type="checkbox" label="Check me out" />
					</Form.Group>
					<Button variant="primary" type="submit">
						Submit
					</Button>
				</Form>
			</Container>
		</section>
	);
};

export default createQuestion;
