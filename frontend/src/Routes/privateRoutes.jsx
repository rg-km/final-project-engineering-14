import React, { useEffect, useContext } from "react";
import { useNavigate, Outlet } from "react-router-dom";
import AuthContext from "../store/auth-context";

function PrivateRoute() {
	const user = useContext(AuthContext);

	const navigate = useNavigate();

	useEffect(() => {
		const token = localStorage.getItem("token");

		if (!user.isLoggedIn && !token) {
			navigate("/login");
		}
		if (user.role === "admin") {
			navigate("/Dashboard");
		}
	}, []);
	return (
		<div>
			<Outlet />
		</div>
	);
}

export default PrivateRoute;
