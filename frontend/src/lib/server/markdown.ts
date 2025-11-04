import { marked } from "marked";
import markedShiki from 'marked-shiki'
import { codeToHtml } from 'shiki'

let isInitialized = false;

// ref: https://shiki.style/guide/best-performance#cache-the-highlighter-instance
const initialiazeMarked = async () => {
	if (isInitialized) {
		return;
	}

	marked.use(
		markedShiki({
			async highlight(code, lang) {
				return await codeToHtml(code, { lang, theme: 'everforest-dark' });
			},
		})
	);

	isInitialized = true;
}

export const getMarked = async () => {
	await initialiazeMarked();
	return marked;
}
