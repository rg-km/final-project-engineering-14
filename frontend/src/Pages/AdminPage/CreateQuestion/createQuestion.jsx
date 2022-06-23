import React from "react";
import "./createQuestion.css";
import Navbar from "../../../Components/Navbar/Navbar";

// React Boostraps
import { Container, Form, Button } from "react-bootstrap";

const createQuestion = () => {
  return (
    <section id="create-question">
      <Navbar />

      <Container className="container-outside">
        <Form onSubmit={""}>
          <h1 className="mb-5">Create Question</h1>
          <Form.Group className="mb-3 mx-auto" controlId="formBasicEmail">
            <Form.Label className="card-title">Programming Language</Form.Label>
            <Form.Control type="text" className="mx-auto form-input" />
          </Form.Group>

          <Form.Group className="mb-3" controlId="formBasicPassword">
            <Form.Label className="card-title">Question</Form.Label>
            <Form.Control as="textarea" className="mx-auto form-area" />
          </Form.Group>
          <Button variant="warning button" className="mt-3" type="submit">
            Submit
          </Button>
        </Form>
      </Container>
    </section>
  );
};

export default createQuestion;
