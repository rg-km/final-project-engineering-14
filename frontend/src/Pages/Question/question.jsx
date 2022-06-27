import React from "react";
import { useNavigate } from "react-router-dom";
import "./question.css";
import { useState, useEffect } from "react";
import API from "../../Api/Api";
import axios from "axios";

//React BS
import Form from "react-bootstrap/Form";

const Question = () => {
	const navigate = useNavigate();
	const [listQuestion, setListQuestion] = useState([]);
	const [listAnswer, setListAnswer] = useState([]);
	const [pageQuestion, setPageQuestion] = useState(1);
	const fetchListPost = async () => {
		let auth = localStorage.getItem("token");
		let pageQuestions = "";
		if (localStorage.getItem("pageQuestion") === null) {
			pageQuestions = pageQuestion;
		} else {
			pageQuestions = localStorage.getItem("pageQuestion");
		}

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
			setListQuestion(res.data.data[0].question);
			setListAnswer(res.data.data[0].answer);
		} catch (error) {}
	};

	useEffect(() => {
		if (localStorage.getItem(pageQuestion) === null) {
			localStorage.setItem("pageQuestion", 1);
			fetchListPost();
		}
		localStorage.removeItem("questionId");
		localStorage.removeItem("Answer Question");
	}, []);

	const handleClick = () => {
		setPageQuestion(pageQuestion + 1);

		localStorage.setItem("pageQuestion", pageQuestion + 1);
		localStorage.setItem("questionId", pageQuestion);

		fetchListPost();
	};

	const handleChange = (e) => {
		e.persist();
		let y = localStorage.getItem("pageQuestion");
		console.log(typeof localStorage.getItem("pageQuestion"));
		let z = localStorage.getItem("questionId");
		let object = {};
		if (y === "1") {
			object[y] = e.target.value;
			localStorage.setItem("Answer Question", JSON.stringify(object));
		} else if (y > z) {
			let x = JSON.parse(localStorage.getItem("Answer Question"));

			x[y] = e.target.value;
			localStorage.setItem("Answer Question", JSON.stringify(x));
		}
	};

	const handleSubmit = async () => {
		let auth = localStorage.getItem("token");
		let answer = localStorage.getItem("Answer Question");
		try {
			let res = await axios.post(
				`${API.API_URL}/home/process-and-result`,
				answer,
				{
					headers: {
						Accept: "/",
						"Content-Type": "application/json",
						Authorization: "Bearer " + auth,
					},
				}
			);

			localStorage.setItem("Level", JSON.stringify(res.data.data));
			navigate("/Level");
		} catch (error) {
			console.log(error);
		}
	};
	let questionId = localStorage.getItem("questionId");
	return (
		<section id="question-page">
			<div className="container">
				<h1>Question</h1>
				<div className="card">
					<div className="card-body">
						<div className="row card-tittle mb-5 mt-3">
							<div className="col-6">
								<h5 className="card-title float-start ml-3">
									Soal ke - {pageQuestion} / 9
								</h5>
							</div>
							<div className="col-6"></div>
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
					{questionId < 8 ? (
						<button
							type="button"
							className="btn btn-warning mr-3"
							onClick={handleClick}
						>
							Next
						</button>
					) : (
						<button
							type="submit"
							className="button-submit btn btn-success mt-3"
							onClick={handleSubmit}
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
