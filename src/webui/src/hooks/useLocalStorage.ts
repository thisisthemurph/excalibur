import { Dispatch, SetStateAction, useCallback, useEffect, useState } from "react";

type SetValue<T> = Dispatch<SetStateAction<T>>;

function useLocalStorage<T>(key: string, initialValue: T): [T, SetValue<T>] {
	const readValue = useCallback((): T => {
		const item = localStorage.getItem(key);
		return item ? (JSON.parse(item) as T) : initialValue;
	}, [initialValue, key]);

	const [storedValue, setStoredValue] = useState<T>(readValue);

	const setValue: SetValue<T> = useCallback((value) => {
		const newValue = value instanceof Function ? value(storedValue) : value;

		localStorage.setItem(key, JSON.stringify(newValue));
		setStoredValue(newValue);
	}, []);

	useEffect(() => {
		const value = readValue();
		localStorage.setItem(key, JSON.stringify(value));
		setStoredValue(value);
	}, []);

	return [storedValue, setValue];
}

export default useLocalStorage;
