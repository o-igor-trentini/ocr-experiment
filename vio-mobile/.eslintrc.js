module.exports = {
	env: {
		'es2021': true,
		'node': true
	},
	extends: [
		'eslint:recommended',
		'plugin:react/recommended',
		'plugin:@typescript-eslint/recommended',
		'plugin:react-hooks/recommended',
		'plugin:@typescript-eslint/eslint-recommended',
		'plugin:@typescript-eslint/recommended-requiring-type-checking'
	],
	parser: '@typescript-eslint/parser',
	parserOptions: {
		'ecmaFeatures': {
			'jsx': true
		},
		'ecmaVersion': 'latest',
		'sourceType': 'module'
	},
	plugins: [
		'react',
		'@typescript-eslint',
		'react-hooks',
		'prettier'
	],
	rules: {
		// 'indent': [
		// 	'error',
		// 	'tab'
		// ],
		'linebreak-style': [
			'error',
			'unix'
		],
		'quotes': [
			'error',
			'single'
		],
		'semi': [
			'error',
			'always'
		],
		'prettier/prettier': 'error',
		'no-empty-function': 'off',
		'@typescript-eslint/no-empty-function': 'off',
		'react/display-name': 'off',
		'react/prop-types': 'off'
	}
};
