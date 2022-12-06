import { Link } from "react-router-dom";

const TemplateHomePage = () => {
  return (
    <>
      <h1>Templates</h1>
      <Link to="/template/create">Create a new template</Link>
    </>
  );
};

export default TemplateHomePage;
