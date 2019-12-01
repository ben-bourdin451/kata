decodeMorse = (morseCode) => {
	return morseCode.trim().split('   ').map(word => {
		return word.split(' ').map(letter => {
			return MORSE_CODE[letter];
		}).join('');
	}).join(' ');
}

Test.describe("Example from description", () => {
	Test.assertEquals(decodeMorse('.... . -.--   .--- ..- -.. .'), 'HEY JUDE')
});
