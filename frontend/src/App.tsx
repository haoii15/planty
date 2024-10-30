import React from "react";
import MyComponent from "./components/data"; // Adjust the import path based on the relative path to the component
import CreateUserForm from "./components/user/create-user-form";

const App: React.FC = () => {
  return (
    <div className="App">
      <CreateUserForm />
    </div>
  );
};

export default App;
