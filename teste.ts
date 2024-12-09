export enum TOKEN_TYPES {
	LEFT_BRACE = "LEFT_BRACE",
	RIGHT_BRACE = "RIGHT_BRACE",
	LEFT_BRACKET = "LEFT_BRACKET",
	RIGHT_BRACKET = "RIGHT_BRACKET",
	COLON = "COLON",
	COMMA = "COMMA",
	STRING = "STRING",
	NUMBER = "NUMBER",
	TRUE = "TRUE",
	FALSE = "FALSE",
	NULL = "NULL",
}

type Token = {
	type: string;
	value: any;
};

export const createToken = (type: TOKEN_TYPES, value?: string): Token => {
	return {
		type,
		value,
	};
};

export const lexer = (input: string): Token[] => {
	let current = 0;
	const tokens: Token[] = [];

	while (current < input.length) {
		let char = input[current];

		if (char === "{") {
			tokens.push(createToken(TOKEN_TYPES.LEFT_BRACE));
			current++;
			continue;
		}

		if (char === "}") {
			tokens.push(createToken(TOKEN_TYPES.RIGHT_BRACE));
			current++;
			continue;
		}

		if (char === "[") {
			tokens.push(createToken(TOKEN_TYPES.LEFT_BRACKET));
			current++;
			continue;
		}

		if (char === "]") {
			tokens.push(createToken(TOKEN_TYPES.RIGHT_BRACKET));
			current++;
			continue;
		}

		if (char === ":") {
			tokens.push(createToken(TOKEN_TYPES.COLON));
			current++;
			continue;
		}

		if (char === ",") {
			tokens.push(createToken(TOKEN_TYPES.COMMA));
			current++;
			continue;
		}

		const WHITESPACE = /\s/;

		if (WHITESPACE.test(char)) {
			current++;
			continue;
		}

		const NUMBERS = /[0-9]/;
		if (NUMBERS.test(char)) {
			let value = "";
			while (NUMBERS.test(char)) {
				value += char;

				char = input[++current];
			}
			tokens.push(createToken(TOKEN_TYPES.NUMBER, value));
			continue;
		}

		if (char === '"') {
			let value = "";
			char = input[++current];
			while (char !== '"') {
				value += char;
				char = input[++current];
			}
			char = input[++current];
			tokens.push(createToken(TOKEN_TYPES.STRING, value));
			continue;
		}

		if (
			char === "t" &&
			input[current + 1] === "r" &&
			input[current + 2] === "u" &&
			input[current + 3] === "e"
		) {
			tokens.push(createToken(TOKEN_TYPES.TRUE));
			current += 4;
			continue;
		}

		if (
			char === "f" &&
			input[current + 1] === "a" &&
			input[current + 2] === "l" &&
			input[current + 3] === "s" &&
			input[current + 4] === "e"
		) {
			tokens.push(createToken(TOKEN_TYPES.FALSE));
			current += 5;
			continue;
		}

		if (
			char === "n" &&
			input[current + 1] === "u" &&
			input[current + 2] === "l" &&
			input[current + 3] === "l"
		) {
			tokens.push(createToken(TOKEN_TYPES.NULL));
			current += 4;
			continue;
		}

		throw new TypeError("I dont know what this character is: " + char);
	}

	return tokens;
};

export const parser = (tokens: Array<{ type: string; value?: any }>) => {
	let current = 0;

	const walk = ():
		| {
				type: string;
				properties?: Array<{ type: string; key: any; value: any }>;
		  }
		| Token => {
		let token = tokens[current];

		if (token.type === TOKEN_TYPES.LEFT_BRACE) {
			token = tokens[++current];

			const node: {
				type: string;
				properties?: Array<{ type: string; key: any; value: any }>;
			} = {
				type: "ObjectExpression",
				properties: [],
			};

			while (token.type !== TOKEN_TYPES.RIGHT_BRACE) {
				const property: { type: string; key: any; value: any } = {
					type: "Property",
					key: token,
					value: null,
				};

				token = tokens[++current];

				token = tokens[++current];
				property.value = walk();
				node.properties?.push(property);

				token = tokens[current];
				if (token.type === TOKEN_TYPES.COMMA) {
					token = tokens[++current];
				}
			}

			current++;
			return node;
		}

		if (token.type === TOKEN_TYPES.RIGHT_BRACE) {
			current++;
			return {
				type: "ObjectExpression",
				properties: [],
			};
		}

		if (token.type === TOKEN_TYPES.LEFT_BRACKET) {
			token = tokens[++current];

			const node: {
				type: string;
				elements?: Array<{ type?: string; value?: any }>;
			} = {
				type: "ArrayExpression",
				elements: [],
			};

			while (token.type !== TOKEN_TYPES.RIGHT_BRACKET) {
				node.elements?.push(walk());
				token = tokens[current];

				if (token.type === TOKEN_TYPES.COMMA) {
					token = tokens[++current];
				}
			}

			current++;
			return node;
		}

		if (token.type === TOKEN_TYPES.STRING) {
			current++;
			return {
				type: "StringLiteral",
				value: token.value,
			};
		}

		if (token.type === TOKEN_TYPES.NUMBER) {
			current++;
			return {
				type: "NumberLiteral",
				value: token.value,
			};
		}

		if (token.type === TOKEN_TYPES.TRUE) {
			current++;
			return {
				type: "BooleanLiteral",
				value: true,
			};
		}

		if (token.type === TOKEN_TYPES.FALSE) {
			current++;
			return {
				type: "BooleanLiteral",
				value: false,
			};
		}

		if (token.type === TOKEN_TYPES.NULL) {
			current++;
			return {
				type: "NullLiteral",
				value: null,
			};
		}

		throw new TypeError(token.type);
	};

	const ast: { type: string; body: any } = {
		type: "Program",
		body: [],
	};

	while (current < tokens.length) {
		ast.body.push(walk());
	}

	return ast;
};

const tokens = lexer(`{
  "key": "value",
  "key-n": 101,
  "key-o": {
    "inner key": "inner value"
  },
  "key-l": ["list value"]
}`);
console.log("tokens", tokens);
const json = parser(tokens);

console.log("parser:", JSON.stringify(json, null, 2));
