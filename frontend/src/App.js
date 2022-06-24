import React from "react";
import { Routes, Route, Navigate } from "react-router-dom";
import {
	Landing,
	Login,
	Register,
	Question,
	Level,
	Language,
	HomePage,
	Dashboard,
	CreateQuestion,
	Recomendation,
} from "./Pages";
import PrivateRoute from "./Routes/privateRoutes";

export const App = () => {
	return (
		<Routes>
			<Route path="/" element={<Landing />}>
				{" "}
			</Route>
			<Route path="/login" element={<Login />}></Route>
			<Route path="/register" element={<Register />}></Route>

			<Route path="" element={<PrivateRoute />}>
				<Route path="/Question" element={<Question />}></Route>
				<Route path="/HomePage" element={<HomePage />}></Route>
				<Route path="/Level" element={<Level />}></Route>
				<Route path="/Language" element={<Language />}></Route>
				<Route path="/Recomendation" element={<Recomendation />}></Route>
				<Route path="/Dashboard" element={<Dashboard />}></Route>
				<Route path="/CreateQuestion" element={<CreateQuestion />}></Route>
			</Route>

			<Route path="*" element={<Navigate to="/" />}></Route>
		</Routes>
	);
};
export default App;
