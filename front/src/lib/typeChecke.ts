export const typeChecker = (tipo: string, data: any, func: Function) => {
	// console.log(tipo, data, func, typeof data)
	if (data && typeof data == tipo) {
		return func(data);
	}
};
