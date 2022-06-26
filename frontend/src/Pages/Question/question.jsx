import React from "react";
import { Link, useParams } from "react-router-dom";
import "./question.css";
import { useState, useEffect } from "react";
import API from "../../Api/Api";
import axios from "axios";

//React BS
import Form from "react-bootstrap/Form";

const Question = () => {
  const [listQuestion, setListQuestion] = useState([]);
  const [listAnswer, setListAnswer] = useState([]);
  const [pageQuestion, setPageQuestion] = useState(1);
  const [questionId, setQuestionId] = useState(0);
  // let { questionId } = useParams();
  // console.log(questionId, "questionId");
  const fetchListPost = async (param) => {
    let auth = localStorage.getItem("token");
    let pageQuestions = "";
    if (localStorage.getItem("pageQuestion") === null) {
      pageQuestions = pageQuestion;
    } else {
      pageQuestions = localStorage.getItem("pageQuestion");
    }

    // e.preventDefault();
    try {
      let res = await axios.get(
        `${API.API_URL}/home/start-ruang-arah/1?page=${pageQuestions}&per_page=1`,
        {
          headers: {
            Accept: "/",
            "Content-Type": "application/json",
            Authorization: "Bearer " + auth,
          },
        }
      );
      console.log(res);
      setListQuestion(res.data.data[0].question);
      setListAnswer(res.data.data[0].answer);
      setQuestionId(res.data.data[0].id);
      console.log(res.data.data);
    } catch (error) {
      console.log(error);
    }
  };

  useEffect(() => {
    console.log(localStorage.getItem("pageQuestion"));
    if (localStorage.getItem(pageQuestion) === null) {
      localStorage.setItem("pageQuestion", 1);
      fetchListPost();
    }
    console.log(questionId);
  }, []);

  const handleClick = () => {
    setPageQuestion(pageQuestion + 1);
    console.log("pageQuestion", pageQuestion);
    localStorage.setItem("pageQuestion", pageQuestion + 1);
    // let y = localStorage.getItem("pageQuestion");
    // console.log("ini y", y);
    fetchListPost();
  };

  const handleChange = (e) => {
    e.persist();
    console.log("test id ", e.target.id);
    let x = pageQuestion;
    let object = {};
    object[pageQuestion] = e.target.value;
    localStorage.setItem("Answer Question", JSON.stringify(object));
    console.log("jawaban", e.target.value);
  };

  return (
    <section id="question-page">
      <div className="container">
        <h1>Question</h1>
        <div className="card">
          <div className="card-body">
            <div className="row card-tittle mb-5 mt-3">
              <div className="col-6">
                <h5 className="card-title float-start ml-3">
                  Soal ke - {pageQuestion} / 10
                </h5>
              </div>
              <div className="col-6">
                {/* <Link to="/Level"> */}
                <button
                  type="button"
                  className="btn btn-warning kanan mr-3"
                  onClick={handleClick}
                >
                  Next
                </button>
                {/* </Link> */}
              </div>
            </div>

            <div className="card-text">
              <p>{listQuestion}</p>
              <Form>
                {["radio"].map((type) => (
                  <div key={`inline-${type}`} className="mb-3">
                    {listAnswer.map((item) => (
                      <Form.Check
                        label={item.answer}
                        id={listQuestion?.id}
                        name="group1"
                        type={type}
                        onChange={handleChange}
                        value={item.answer}
                      />
                    ))}
                  </div>
                ))}
              </Form>
            </div>
          </div>
          {pageQuestion === 10 ? (
            <button
              type="submit"
              className="button-submit btn btn-success mt-3"
            >
              Submit
            </button>
          ) : (
            <button
              type="submit"
              className="button-submit btn btn-secondary mt-3"
              disabled
            >
              Submit
            </button>
          )}
        </div>
      </div>
    </section>
  );
};

export default Question;
