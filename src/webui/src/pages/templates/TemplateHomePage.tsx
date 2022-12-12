import { Link } from "react-router-dom";

const TemplateHomePage = () => {
  return (
    <>
      <h1 className="px-wrap py-wrap">Templates</h1>
      <main className="px-wrap">
        <Link to="/template/create">Create a new template</Link>
      </main>
    </>
  );
};

export default TemplateHomePage;
