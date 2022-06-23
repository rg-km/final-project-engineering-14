import React, { useState } from "react";

const AuthContext = React.createContext({
	token: "",
	isLoggedIn: false,
	login: (token) => {},
	logout: () => {},
});

export const AuthContextProvider = (props) => {
	const [token, setToken] = useState(null);
	const [role, setRole] = useState(null);
	const [isLoggedIn, setIsLoggedIn] = useState(false);

	const loginHandler = (token, role) => {
		setToken(token);
		setRole(role);
		setIsLoggedIn(true);
	};

	const logoutHandler = () => {
		setToken(null);
	};

	const contextValue = {
		token: token,
		role: role,
		isLoggedIn: isLoggedIn,
		login: loginHandler,
		logout: logoutHandler,
	};

	return (
		<AuthContext.Provider value={contextValue}>
			{props.children}
		</AuthContext.Provider>
	);
};

export default AuthContext;
