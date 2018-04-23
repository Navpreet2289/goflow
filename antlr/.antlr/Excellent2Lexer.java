// Generated from /Users/rowan/Nyaruka/go/src/github.com/nyaruka/goflow/antlr/Excellent2.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.Lexer;
import org.antlr.v4.runtime.CharStream;
import org.antlr.v4.runtime.Token;
import org.antlr.v4.runtime.TokenStream;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.misc.*;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Excellent2Lexer extends Lexer {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		COMMA=1, LPAREN=2, RPAREN=3, LBRACK=4, RBRACK=5, DOT=6, PLUS=7, MINUS=8, 
		TIMES=9, DIVIDE=10, EXPONENT=11, EQ=12, NEQ=13, LTE=14, LT=15, GTE=16, 
		GT=17, AMPERSAND=18, DECIMAL=19, STRING=20, TRUE=21, FALSE=22, NULL=23, 
		NAME=24, WS=25, ERROR=26;
	public static String[] channelNames = {
		"DEFAULT_TOKEN_CHANNEL", "HIDDEN"
	};

	public static String[] modeNames = {
		"DEFAULT_MODE"
	};

	public static final String[] ruleNames = {
		"COMMA", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "PLUS", "MINUS", 
		"TIMES", "DIVIDE", "EXPONENT", "EQ", "NEQ", "LTE", "LT", "GTE", "GT", 
		"AMPERSAND", "DECIMAL", "STRING", "TRUE", "FALSE", "NULL", "NAME", "WS", 
		"ERROR", "UnicodeLetter", "UnicodeClass_LU", "UnicodeClass_LL", "UnicodeClass_LT", 
		"UnicodeClass_LM", "UnicodeClass_LO", "UnicodeDigit"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "','", "'('", "')'", "'['", "']'", "'.'", "'+'", "'-'", "'*'", "'/'", 
		"'^'", "'='", "'!='", "'<='", "'<'", "'>='", "'>'", "'&'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, "COMMA", "LPAREN", "RPAREN", "LBRACK", "RBRACK", "DOT", "PLUS", 
		"MINUS", "TIMES", "DIVIDE", "EXPONENT", "EQ", "NEQ", "LTE", "LT", "GTE", 
		"GT", "AMPERSAND", "DECIMAL", "STRING", "TRUE", "FALSE", "NULL", "NAME", 
		"WS", "ERROR"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}


	public Excellent2Lexer(CharStream input) {
		super(input);
		_interp = new LexerATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@Override
	public String getGrammarFileName() { return "Excellent2.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public String[] getChannelNames() { return channelNames; }

	@Override
	public String[] getModeNames() { return modeNames; }

	@Override
	public ATN getATN() { return _ATN; }

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\2\34\u00bd\b\1\4\2"+
		"\t\2\4\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4"+
		"\13\t\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22"+
		"\t\22\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31"+
		"\t\31\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t"+
		" \4!\t!\4\"\t\"\3\2\3\2\3\3\3\3\3\4\3\4\3\5\3\5\3\6\3\6\3\7\3\7\3\b\3"+
		"\b\3\t\3\t\3\n\3\n\3\13\3\13\3\f\3\f\3\r\3\r\3\16\3\16\3\16\3\17\3\17"+
		"\3\17\3\20\3\20\3\21\3\21\3\21\3\22\3\22\3\23\3\23\3\24\6\24n\n\24\r\24"+
		"\16\24o\3\24\3\24\6\24t\n\24\r\24\16\24u\5\24x\n\24\3\25\3\25\3\25\3\25"+
		"\7\25~\n\25\f\25\16\25\u0081\13\25\3\25\3\25\3\26\3\26\3\26\3\26\3\26"+
		"\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\31\6\31\u0096"+
		"\n\31\r\31\16\31\u0097\3\31\3\31\3\31\7\31\u009d\n\31\f\31\16\31\u00a0"+
		"\13\31\3\32\6\32\u00a3\n\32\r\32\16\32\u00a4\3\32\3\32\3\33\3\33\3\34"+
		"\3\34\3\34\3\34\3\34\5\34\u00b0\n\34\3\35\3\35\3\36\3\36\3\37\3\37\3 "+
		"\3 \3!\3!\3\"\3\"\2\2#\3\3\5\4\7\5\t\6\13\7\r\b\17\t\21\n\23\13\25\f\27"+
		"\r\31\16\33\17\35\20\37\21!\22#\23%\24\'\25)\26+\27-\30/\31\61\32\63\33"+
		"\65\34\67\29\2;\2=\2?\2A\2C\2\3\2\24\3\2\62;\3\2$$\4\2VVvv\4\2TTtt\4\2"+
		"WWww\4\2GGgg\4\2HHhh\4\2CCcc\4\2NNnn\4\2UUuu\4\2PPpp\5\2\13\f\17\17\""+
		"\"T\2C\\\u00c2\u00d8\u00da\u00e0\u0102\u0138\u013b\u0149\u014c\u017f\u0183"+
		"\u0184\u0186\u018d\u0190\u0193\u0195\u0196\u0198\u019a\u019e\u019f\u01a1"+
		"\u01a2\u01a4\u01ab\u01ae\u01b5\u01b7\u01be\u01c6\u01cf\u01d1\u01dd\u01e0"+
		"\u01f0\u01f3\u01f6\u01f8\u01fa\u01fc\u0234\u023c\u023d\u023f\u0240\u0243"+
		"\u0248\u024a\u0250\u0372\u0374\u0378\u0381\u0388\u038c\u038e\u03a3\u03a5"+
		"\u03ad\u03d1\u03d6\u03da\u03f0\u03f6\u03f9\u03fb\u03fc\u03ff\u0431\u0462"+
		"\u0482\u048c\u04cf\u04d2\u0530\u0533\u0558\u10a2\u10c7\u10c9\u10cf\u1e02"+
		"\u1e96\u1ea0\u1f00\u1f0a\u1f11\u1f1a\u1f1f\u1f2a\u1f31\u1f3a\u1f41\u1f4a"+
		"\u1f4f\u1f5b\u1f61\u1f6a\u1f71\u1fba\u1fbd\u1fca\u1fcd\u1fda\u1fdd\u1fea"+
		"\u1fee\u1ffa\u1ffd\u2104\u2109\u210d\u210f\u2112\u2114\u2117\u211f\u2126"+
		"\u212f\u2132\u2135\u2140\u2141\u2147\u2185\u2c02\u2c30\u2c62\u2c66\u2c69"+
		"\u2c72\u2c74\u2c77\u2c80\u2c82\u2c84\u2ce4\u2ced\u2cef\u2cf4\ua642\ua644"+
		"\ua66e\ua682\ua69c\ua724\ua730\ua734\ua770\ua77b\ua788\ua78d\ua78f\ua792"+
		"\ua794\ua798\ua7af\ua7b2\ua7b3\uff23\uff3cS\2c|\u00b7\u00f8\u00fa\u0101"+
		"\u0103\u0179\u017c\u0182\u0185\u0187\u018a\u0194\u0197\u019d\u01a0\u01a3"+
		"\u01a5\u01a7\u01aa\u01af\u01b2\u01b6\u01b8\u01c1\u01c8\u01ce\u01d0\u01f5"+
		"\u01f7\u01fb\u01fd\u023b\u023e\u0244\u0249\u0295\u0297\u02b1\u0373\u0375"+
		"\u0379\u037f\u0392\u03d0\u03d2\u03d3\u03d7\u03d9\u03db\u03f5\u03f7\u0461"+
		"\u0463\u0483\u048d\u04c1\u04c4\u0531\u0563\u0589\u1d02\u1d2d\u1d6d\u1d79"+
		"\u1d7b\u1d9c\u1e03\u1e9f\u1ea1\u1f09\u1f12\u1f17\u1f22\u1f29\u1f32\u1f39"+
		"\u1f42\u1f47\u1f52\u1f59\u1f62\u1f69\u1f72\u1f7f\u1f82\u1f89\u1f92\u1f99"+
		"\u1fa2\u1fa9\u1fb2\u1fb6\u1fb8\u1fb9\u1fc0\u1fc6\u1fc8\u1fc9\u1fd2\u1fd5"+
		"\u1fd8\u1fd9\u1fe2\u1fe9\u1ff4\u1ff6\u1ff8\u1ff9\u210c\u2115\u2131\u213b"+
		"\u213e\u213f\u2148\u214b\u2150\u2186\u2c32\u2c60\u2c63\u2c6e\u2c73\u2c7d"+
		"\u2c83\u2cee\u2cf0\u2cf5\u2d02\u2d27\u2d29\u2d2f\ua643\ua66f\ua683\ua69d"+
		"\ua725\ua733\ua735\ua77a\ua77c\ua77e\ua781\ua789\ua78e\ua790\ua793\ua797"+
		"\ua799\ua7ab\ua7fc\uab5c\uab66\uab67\ufb02\ufb08\ufb15\ufb19\uff43\uff5c"+
		"\b\2\u01c7\u01cd\u01f4\u1f91\u1f9a\u1fa1\u1faa\u1fb1\u1fbe\u1fce\u1ffe"+
		"\u1ffe#\2\u02b2\u02c3\u02c8\u02d3\u02e2\u02e6\u02ee\u02f0\u0376\u037c"+
		"\u055b\u0642\u06e7\u06e8\u07f6\u07f7\u07fc\u081c\u0826\u082a\u0973\u0e48"+
		"\u0ec8\u10fe\u17d9\u1845\u1aa9\u1c7f\u1d2e\u1d6c\u1d7a\u1dc1\u2073\u2081"+
		"\u2092\u209e\u2c7e\u2c7f\u2d71\u2e31\u3007\u3037\u303d\u3100\ua017\ua4ff"+
		"\ua60e\ua681\ua69e\ua69f\ua719\ua721\ua772\ua78a\ua7fa\ua7fb\ua9d1\ua9e8"+
		"\uaa72\uaadf\uaaf5\uaaf6\uab5e\uab61\uff72\uffa1\u00ec\2\u00ac\u00bc\u01bd"+
		"\u01c5\u0296\u05ec\u05f2\u05f4\u0622\u0641\u0643\u064c\u0670\u0671\u0673"+
		"\u06d5\u06d7\u06fe\u0701\u0712\u0714\u0731\u074f\u07a7\u07b3\u07ec\u0802"+
		"\u0817\u0842\u085a\u08a2\u08b4\u0906\u093b\u093f\u0952\u095a\u0963\u0974"+
		"\u0982\u0987\u098e\u0991\u0992\u0995\u09aa\u09ac\u09b2\u09b4\u09bb\u09bf"+
		"\u09d0\u09de\u09df\u09e1\u09e3\u09f2\u09f3\u0a07\u0a0c\u0a11\u0a12\u0a15"+
		"\u0a2a\u0a2c\u0a32\u0a34\u0a35\u0a37\u0a38\u0a3a\u0a3b\u0a5b\u0a5e\u0a60"+
		"\u0a76\u0a87\u0a8f\u0a91\u0a93\u0a95\u0aaa\u0aac\u0ab2\u0ab4\u0ab5\u0ab7"+
		"\u0abb\u0abf\u0ad2\u0ae2\u0ae3\u0b07\u0b0e\u0b11\u0b12\u0b15\u0b2a\u0b2c"+
		"\u0b32\u0b34\u0b35\u0b37\u0b3b\u0b3f\u0b63\u0b73\u0b85\u0b87\u0b8c\u0b90"+
		"\u0b92\u0b94\u0b97\u0b9b\u0b9c\u0b9e\u0bac\u0bb0\u0bbb\u0bd2\u0c0e\u0c10"+
		"\u0c12\u0c14\u0c2a\u0c2c\u0c3b\u0c3f\u0c8e\u0c90\u0c92\u0c94\u0caa\u0cac"+
		"\u0cb5\u0cb7\u0cbb\u0cbf\u0ce0\u0ce2\u0ce3\u0cf3\u0cf4\u0d07\u0d0e\u0d10"+
		"\u0d12\u0d14\u0d3c\u0d3f\u0d50\u0d62\u0d63\u0d7c\u0d81\u0d87\u0d98\u0d9c"+
		"\u0db3\u0db5\u0dbd\u0dbf\u0dc8\u0e03\u0e32\u0e34\u0e35\u0e42\u0e47\u0e83"+
		"\u0e84\u0e86\u0e8c\u0e8f\u0e99\u0e9b\u0ea1\u0ea3\u0ea5\u0ea7\u0ea9\u0eac"+
		"\u0ead\u0eaf\u0eb2\u0eb4\u0eb5\u0ebf\u0ec6\u0ede\u0ee1\u0f02\u0f49\u0f4b"+
		"\u0f6e\u0f8a\u0f8e\u1002\u102c\u1041\u1057\u105c\u105f\u1063\u1072\u1077"+
		"\u1083\u1090\u10fc\u10ff\u124a\u124c\u124f\u1252\u1258\u125a\u125f\u1262"+
		"\u128a\u128c\u128f\u1292\u12b2\u12b4\u12b7\u12ba\u12c0\u12c2\u12c7\u12ca"+
		"\u12d8\u12da\u1312\u1314\u1317\u131a\u135c\u1382\u1391\u13a2\u13f6\u1403"+
		"\u166e\u1671\u1681\u1683\u169c\u16a2\u16ec\u16f3\u16fa\u1702\u170e\u1710"+
		"\u1713\u1722\u1733\u1742\u1753\u1762\u176e\u1770\u1772\u1782\u17b5\u17de"+
		"\u1844\u1846\u1879\u1882\u18aa\u18ac\u18f7\u1902\u1920\u1952\u196f\u1972"+
		"\u1976\u1982\u19ad\u19c3\u19c9\u1a02\u1a18\u1a22\u1a56\u1b07\u1b35\u1b47"+
		"\u1b4d\u1b85\u1ba2\u1bb0\u1bb1\u1bbc\u1be7\u1c02\u1c25\u1c4f\u1c51\u1c5c"+
		"\u1c79\u1ceb\u1cee\u1cf0\u1cf3\u1cf7\u1cf8\u2137\u213a\u2d32\u2d69\u2d82"+
		"\u2d98\u2da2\u2da8\u2daa\u2db0\u2db2\u2db8\u2dba\u2dc0\u2dc2\u2dc8\u2dca"+
		"\u2dd0\u2dd2\u2dd8\u2dda\u2de0\u3008\u303e\u3043\u3098\u30a1\u30fc\u3101"+
		"\u312f\u3133\u3190\u31a2\u31bc\u31f2\u3201\u3402\u4db7\u4e02\u9fce\ua002"+
		"\ua016\ua018\ua48e\ua4d2\ua4f9\ua502\ua60d\ua612\ua621\ua62c\ua62d\ua670"+
		"\ua6e7\ua7f9\ua803\ua805\ua807\ua809\ua80c\ua80e\ua824\ua842\ua875\ua884"+
		"\ua8b5\ua8f4\ua8f9\ua8fd\ua927\ua932\ua948\ua962\ua97e\ua986\ua9b4\ua9e2"+
		"\ua9e6\ua9e9\ua9f1\ua9fc\uaa00\uaa02\uaa2a\uaa42\uaa44\uaa46\uaa4d\uaa62"+
		"\uaa71\uaa73\uaa78\uaa7c\uaab1\uaab3\uaabf\uaac2\uaac4\uaadd\uaade\uaae2"+
		"\uaaec\uaaf4\uab08\uab0b\uab10\uab13\uab18\uab22\uab28\uab2a\uab30\uabc2"+
		"\uabe4\uac02\ud7a5\ud7b2\ud7c8\ud7cd\ud7fd\uf902\ufa6f\ufa72\ufadb\ufb1f"+
		"\ufb2a\ufb2c\ufb38\ufb3a\ufb3e\ufb40\ufbb3\ufbd5\ufd3f\ufd52\ufd91\ufd94"+
		"\ufdc9\ufdf2\ufdfd\ufe72\ufe76\ufe78\ufefe\uff68\uff71\uff73\uff9f\uffa2"+
		"\uffc0\uffc4\uffc9\uffcc\uffd1\uffd4\uffd9\uffdc\uffde\'\2\62;\u0662\u066b"+
		"\u06f2\u06fb\u07c2\u07cb\u0968\u0971\u09e8\u09f1\u0a68\u0a71\u0ae8\u0af1"+
		"\u0b68\u0b71\u0be8\u0bf1\u0c68\u0c71\u0ce8\u0cf1\u0d68\u0d71\u0de8\u0df1"+
		"\u0e52\u0e5b\u0ed2\u0edb\u0f22\u0f2b\u1042\u104b\u1092\u109b\u17e2\u17eb"+
		"\u1812\u181b\u1948\u1951\u19d2\u19db\u1a82\u1a8b\u1a92\u1a9b\u1b52\u1b5b"+
		"\u1bb2\u1bbb\u1c42\u1c4b\u1c52\u1c5b\ua622\ua62b\ua8d2\ua8db\ua902\ua90b"+
		"\ua9d2\ua9db\ua9f2\ua9fb\uaa52\uaa5b\uabf2\uabfb\uff12\uff1b\2\u00c3\2"+
		"\3\3\2\2\2\2\5\3\2\2\2\2\7\3\2\2\2\2\t\3\2\2\2\2\13\3\2\2\2\2\r\3\2\2"+
		"\2\2\17\3\2\2\2\2\21\3\2\2\2\2\23\3\2\2\2\2\25\3\2\2\2\2\27\3\2\2\2\2"+
		"\31\3\2\2\2\2\33\3\2\2\2\2\35\3\2\2\2\2\37\3\2\2\2\2!\3\2\2\2\2#\3\2\2"+
		"\2\2%\3\2\2\2\2\'\3\2\2\2\2)\3\2\2\2\2+\3\2\2\2\2-\3\2\2\2\2/\3\2\2\2"+
		"\2\61\3\2\2\2\2\63\3\2\2\2\2\65\3\2\2\2\3E\3\2\2\2\5G\3\2\2\2\7I\3\2\2"+
		"\2\tK\3\2\2\2\13M\3\2\2\2\rO\3\2\2\2\17Q\3\2\2\2\21S\3\2\2\2\23U\3\2\2"+
		"\2\25W\3\2\2\2\27Y\3\2\2\2\31[\3\2\2\2\33]\3\2\2\2\35`\3\2\2\2\37c\3\2"+
		"\2\2!e\3\2\2\2#h\3\2\2\2%j\3\2\2\2\'m\3\2\2\2)y\3\2\2\2+\u0084\3\2\2\2"+
		"-\u0089\3\2\2\2/\u008f\3\2\2\2\61\u0095\3\2\2\2\63\u00a2\3\2\2\2\65\u00a8"+
		"\3\2\2\2\67\u00af\3\2\2\29\u00b1\3\2\2\2;\u00b3\3\2\2\2=\u00b5\3\2\2\2"+
		"?\u00b7\3\2\2\2A\u00b9\3\2\2\2C\u00bb\3\2\2\2EF\7.\2\2F\4\3\2\2\2GH\7"+
		"*\2\2H\6\3\2\2\2IJ\7+\2\2J\b\3\2\2\2KL\7]\2\2L\n\3\2\2\2MN\7_\2\2N\f\3"+
		"\2\2\2OP\7\60\2\2P\16\3\2\2\2QR\7-\2\2R\20\3\2\2\2ST\7/\2\2T\22\3\2\2"+
		"\2UV\7,\2\2V\24\3\2\2\2WX\7\61\2\2X\26\3\2\2\2YZ\7`\2\2Z\30\3\2\2\2[\\"+
		"\7?\2\2\\\32\3\2\2\2]^\7#\2\2^_\7?\2\2_\34\3\2\2\2`a\7>\2\2ab\7?\2\2b"+
		"\36\3\2\2\2cd\7>\2\2d \3\2\2\2ef\7@\2\2fg\7?\2\2g\"\3\2\2\2hi\7@\2\2i"+
		"$\3\2\2\2jk\7(\2\2k&\3\2\2\2ln\t\2\2\2ml\3\2\2\2no\3\2\2\2om\3\2\2\2o"+
		"p\3\2\2\2pw\3\2\2\2qs\7\60\2\2rt\t\2\2\2sr\3\2\2\2tu\3\2\2\2us\3\2\2\2"+
		"uv\3\2\2\2vx\3\2\2\2wq\3\2\2\2wx\3\2\2\2x(\3\2\2\2y\177\7$\2\2z~\n\3\2"+
		"\2{|\7$\2\2|~\7$\2\2}z\3\2\2\2}{\3\2\2\2~\u0081\3\2\2\2\177}\3\2\2\2\177"+
		"\u0080\3\2\2\2\u0080\u0082\3\2\2\2\u0081\177\3\2\2\2\u0082\u0083\7$\2"+
		"\2\u0083*\3\2\2\2\u0084\u0085\t\4\2\2\u0085\u0086\t\5\2\2\u0086\u0087"+
		"\t\6\2\2\u0087\u0088\t\7\2\2\u0088,\3\2\2\2\u0089\u008a\t\b\2\2\u008a"+
		"\u008b\t\t\2\2\u008b\u008c\t\n\2\2\u008c\u008d\t\13\2\2\u008d\u008e\t"+
		"\7\2\2\u008e.\3\2\2\2\u008f\u0090\t\f\2\2\u0090\u0091\t\6\2\2\u0091\u0092"+
		"\t\n\2\2\u0092\u0093\t\n\2\2\u0093\60\3\2\2\2\u0094\u0096\5\67\34\2\u0095"+
		"\u0094\3\2\2\2\u0096\u0097\3\2\2\2\u0097\u0095\3\2\2\2\u0097\u0098\3\2"+
		"\2\2\u0098\u009e\3\2\2\2\u0099\u009d\5\67\34\2\u009a\u009d\5C\"\2\u009b"+
		"\u009d\7a\2\2\u009c\u0099\3\2\2\2\u009c\u009a\3\2\2\2\u009c\u009b\3\2"+
		"\2\2\u009d\u00a0\3\2\2\2\u009e\u009c\3\2\2\2\u009e\u009f\3\2\2\2\u009f"+
		"\62\3\2\2\2\u00a0\u009e\3\2\2\2\u00a1\u00a3\t\r\2\2\u00a2\u00a1\3\2\2"+
		"\2\u00a3\u00a4\3\2\2\2\u00a4\u00a2\3\2\2\2\u00a4\u00a5\3\2\2\2\u00a5\u00a6"+
		"\3\2\2\2\u00a6\u00a7\b\32\2\2\u00a7\64\3\2\2\2\u00a8\u00a9\13\2\2\2\u00a9"+
		"\66\3\2\2\2\u00aa\u00b0\59\35\2\u00ab\u00b0\5;\36\2\u00ac\u00b0\5=\37"+
		"\2\u00ad\u00b0\5? \2\u00ae\u00b0\5A!\2\u00af\u00aa\3\2\2\2\u00af\u00ab"+
		"\3\2\2\2\u00af\u00ac\3\2\2\2\u00af\u00ad\3\2\2\2\u00af\u00ae\3\2\2\2\u00b0"+
		"8\3\2\2\2\u00b1\u00b2\t\16\2\2\u00b2:\3\2\2\2\u00b3\u00b4\t\17\2\2\u00b4"+
		"<\3\2\2\2\u00b5\u00b6\t\20\2\2\u00b6>\3\2\2\2\u00b7\u00b8\t\21\2\2\u00b8"+
		"@\3\2\2\2\u00b9\u00ba\t\22\2\2\u00baB\3\2\2\2\u00bb\u00bc\t\23\2\2\u00bc"+
		"D\3\2\2\2\r\2ouw}\177\u0097\u009c\u009e\u00a4\u00af\3\b\2\2";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}