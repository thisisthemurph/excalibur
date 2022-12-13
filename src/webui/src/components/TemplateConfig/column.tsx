import { FormColumnSchemaType } from "./z";

interface Props {
	index: number;
	deleteColumn: (index: number) => void;
	updateColumn: (index: number, updated: FormColumnSchemaType) => void;
}

enum DataType {
	String = "String",
	Number = "Number",
	Boolean = "Boolean",
}

const TemplateConfigColumn = ({ index, deleteColumn, updateColumn }: Props) => {
	const handleDeleteColumn = () => {
		deleteColumn(index);
	};

	return (
		<form className="column">
			<section className="column__data">
				<label htmlFor="col__datatype--0">Data type</label>
				<select id="col__datatype--0">
					<option value={DataType.String}>{DataType.String}</option>
					<option value={DataType.Number}>{DataType.Number}</option>
					<option value={DataType.Boolean}>{DataType.Boolean}</option>
				</select>
			</section>

			<section className="column__data">
				<label htmlFor="col__actualName--0">Actual name</label>
				<input type="text" id="col__actualName--0" />
			</section>

			<section className="column__data">
				<label htmlFor="col__prettyName--0">Pretty name</label>
				<input type="text" id="col__prettyName--0" />
			</section>

			<section className="column__data">
				<button onClick={handleDeleteColumn}>Delete</button>
			</section>
		</form>
	);
};

export { TemplateConfigColumn };
