// Generated from Chat.g4 by ANTLR 4.8
// jshint ignore: start
var antlr4 = require('antlr4/index');
var ChatListener = require('./ChatListener').ChatListener;
var grammarFileName = "Chat.g4";


var serializedATN = ["\u0003\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964",
    "\u0003\u0014p\u0004\u0002\t\u0002\u0004\u0003\t\u0003\u0004\u0004\t",
    "\u0004\u0004\u0005\t\u0005\u0004\u0006\t\u0006\u0004\u0007\t\u0007\u0004",
    "\b\t\b\u0004\t\t\t\u0004\n\t\n\u0004\u000b\t\u000b\u0004\f\t\f\u0004",
    "\r\t\r\u0004\u000e\t\u000e\u0004\u000f\t\u000f\u0003\u0002\u0006\u0002",
    " \n\u0002\r\u0002\u000e\u0002!\u0003\u0002\u0003\u0002\u0003\u0003\u0003",
    "\u0003\u0003\u0003\u0003\u0003\u0003\u0004\u0003\u0004\u0003\u0004\u0003",
    "\u0004\u0003\u0004\u0006\u0004/\n\u0004\r\u0004\u000e\u00040\u0003\u0005",
    "\u0003\u0005\u0003\u0006\u0003\u0006\u0003\u0006\u0003\u0006\u0003\u0006",
    "\u0005\u0006:\n\u0006\u0003\u0006\u0005\u0006=\n\u0006\u0003\u0006\u0003",
    "\u0006\u0005\u0006A\n\u0006\u0003\u0006\u0003\u0006\u0005\u0006E\n\u0006",
    "\u0003\u0006\u0003\u0006\u0005\u0006I\n\u0006\u0005\u0006K\n\u0006\u0003",
    "\u0007\u0003\u0007\u0003\b\u0003\b\u0003\t\u0003\t\u0003\n\u0003\n\u0003",
    "\u000b\u0003\u000b\u0003\u000b\u0003\f\u0003\f\u0005\fZ\n\f\u0003\f",
    "\u0003\f\u0003\f\u0005\f_\n\f\u0003\f\u0005\fb\n\f\u0003\r\u0003\r\u0003",
    "\r\u0003\u000e\u0003\u000e\u0003\u000e\u0003\u000e\u0003\u000e\u0003",
    "\u000e\u0003\u000f\u0003\u000f\u0003\u000f\u0003\u000f\u0002\u0002\u0010",
    "\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c",
    "\u0002\u0003\u0003\u0002\u000f\u0010\u0002p\u0002\u001f\u0003\u0002",
    "\u0002\u0002\u0004%\u0003\u0002\u0002\u0002\u0006.\u0003\u0002\u0002",
    "\u0002\b2\u0003\u0002\u0002\u0002\nJ\u0003\u0002\u0002\u0002\fL\u0003",
    "\u0002\u0002\u0002\u000eN\u0003\u0002\u0002\u0002\u0010P\u0003\u0002",
    "\u0002\u0002\u0012R\u0003\u0002\u0002\u0002\u0014T\u0003\u0002\u0002",
    "\u0002\u0016a\u0003\u0002\u0002\u0002\u0018c\u0003\u0002\u0002\u0002",
    "\u001af\u0003\u0002\u0002\u0002\u001cl\u0003\u0002\u0002\u0002\u001e",
    " \u0005\u0004\u0003\u0002\u001f\u001e\u0003\u0002\u0002\u0002 !\u0003",
    "\u0002\u0002\u0002!\u001f\u0003\u0002\u0002\u0002!\"\u0003\u0002\u0002",
    "\u0002\"#\u0003\u0002\u0002\u0002#$\u0007\u0002\u0002\u0003$\u0003\u0003",
    "\u0002\u0002\u0002%&\u0005\b\u0005\u0002&\'\u0005\n\u0006\u0002\'(\u0007",
    "\u0011\u0002\u0002(\u0005\u0003\u0002\u0002\u0002)/\u0005\u0016\f\u0002",
    "*/\u0005\u0018\r\u0002+/\u0005\u001a\u000e\u0002,/\u0005\u001c\u000f",
    "\u0002-/\u0007\u000e\u0002\u0002.)\u0003\u0002\u0002\u0002.*\u0003\u0002",
    "\u0002\u0002.+\u0003\u0002\u0002\u0002.,\u0003\u0002\u0002\u0002.-\u0003",
    "\u0002\u0002\u0002/0\u0003\u0002\u0002\u00020.\u0003\u0002\u0002\u0002",
    "01\u0003\u0002\u0002\u00021\u0007\u0003\u0002\u0002\u000223\u0007\u000e",
    "\u0002\u00023\t\u0003\u0002\u0002\u000245\t\u0002\u0002\u000256\u0007",
    "\u0003\u0002\u00026K\u0005\u0006\u0004\u000279\u0005\f\u0007\u00028",
    ":\u0007\u0004\u0002\u000298\u0003\u0002\u0002\u00029:\u0003\u0002\u0002",
    "\u0002:<\u0003\u0002\u0002\u0002;=\u0005\u000e\b\u0002<;\u0003\u0002",
    "\u0002\u0002<=\u0003\u0002\u0002\u0002=@\u0003\u0002\u0002\u0002>?\u0007",
    "\u0005\u0002\u0002?A\u0005\u0010\t\u0002@>\u0003\u0002\u0002\u0002@",
    "A\u0003\u0002\u0002\u0002AD\u0003\u0002\u0002\u0002BC\u0007\u0006\u0002",
    "\u0002CE\u0005\u0012\n\u0002DB\u0003\u0002\u0002\u0002DE\u0003\u0002",
    "\u0002\u0002EH\u0003\u0002\u0002\u0002FG\u0007\u0007\u0002\u0002GI\u0005",
    "\u0014\u000b\u0002HF\u0003\u0002\u0002\u0002HI\u0003\u0002\u0002\u0002",
    "IK\u0003\u0002\u0002\u0002J4\u0003\u0002\u0002\u0002J7\u0003\u0002\u0002",
    "\u0002K\u000b\u0003\u0002\u0002\u0002LM\u0007\u000e\u0002\u0002M\r\u0003",
    "\u0002\u0002\u0002NO\u0007\u000e\u0002\u0002O\u000f\u0003\u0002\u0002",
    "\u0002PQ\u0005\u0006\u0004\u0002Q\u0011\u0003\u0002\u0002\u0002RS\u0005",
    "\u0006\u0004\u0002S\u0013\u0003\u0002\u0002\u0002TU\u0007\u0013\u0002",
    "\u0002UV\u0007\r\u0002\u0002V\u0015\u0003\u0002\u0002\u0002WY\u0007",
    "\u0003\u0002\u0002XZ\u0007\b\u0002\u0002YX\u0003\u0002\u0002\u0002Y",
    "Z\u0003\u0002\u0002\u0002Z[\u0003\u0002\u0002\u0002[b\u0007\t\u0002",
    "\u0002\\^\u0007\u0003\u0002\u0002]_\u0007\b\u0002\u0002^]\u0003\u0002",
    "\u0002\u0002^_\u0003\u0002\u0002\u0002_`\u0003\u0002\u0002\u0002`b\u0007",
    "\n\u0002\u0002aW\u0003\u0002\u0002\u0002a\\\u0003\u0002\u0002\u0002",
    "b\u0017\u0003\u0002\u0002\u0002cd\u0007\u0012\u0002\u0002de\u0007\u0012",
    "\u0002\u0002e\u0019\u0003\u0002\u0002\u0002fg\u0007\u000b\u0002\u0002",
    "gh\u0007\u000e\u0002\u0002hi\u0007\u000b\u0002\u0002ij\u0005\u0006\u0004",
    "\u0002jk\u0007\u000b\u0002\u0002k\u001b\u0003\u0002\u0002\u0002lm\u0007",
    "\f\u0002\u0002mn\u0007\u000e\u0002\u0002n\u001d\u0003\u0002\u0002\u0002",
    "\u000e!.09<@DHJY^a"].join("");


var atn = new antlr4.atn.ATNDeserializer().deserialize(serializedATN);

var decisionsToDFA = atn.decisionToState.map( function(ds, index) { return new antlr4.dfa.DFA(ds, index); });

var sharedContextCache = new antlr4.PredictionContextCache();

var literalNames = [ null, "':'", "'to'", "'because'", "'by'", "'for'", 
                     "'-'", "')'", "'('", "'/'", "'@'" ];

var symbolicNames = [ null, null, null, null, null, null, null, null, null, 
                      null, null, "TIME_UNIT", "IDENT", "SAYS", "SHOUTS", 
                      "NEWLINE", "TEXT", "NUMBER", "WS" ];

var ruleNames =  [ "chat", "line", "message", "name", "command", "verb", 
                   "object", "reason", "method", "time", "emoticon", "link", 
                   "color", "mention" ];

function ChatParser (input) {
	antlr4.Parser.call(this, input);
    this._interp = new antlr4.atn.ParserATNSimulator(this, atn, decisionsToDFA, sharedContextCache);
    this.ruleNames = ruleNames;
    this.literalNames = literalNames;
    this.symbolicNames = symbolicNames;
    return this;
}

ChatParser.prototype = Object.create(antlr4.Parser.prototype);
ChatParser.prototype.constructor = ChatParser;

Object.defineProperty(ChatParser.prototype, "atn", {
	get : function() {
		return atn;
	}
});

ChatParser.EOF = antlr4.Token.EOF;
ChatParser.T__0 = 1;
ChatParser.T__1 = 2;
ChatParser.T__2 = 3;
ChatParser.T__3 = 4;
ChatParser.T__4 = 5;
ChatParser.T__5 = 6;
ChatParser.T__6 = 7;
ChatParser.T__7 = 8;
ChatParser.T__8 = 9;
ChatParser.T__9 = 10;
ChatParser.TIME_UNIT = 11;
ChatParser.IDENT = 12;
ChatParser.SAYS = 13;
ChatParser.SHOUTS = 14;
ChatParser.NEWLINE = 15;
ChatParser.TEXT = 16;
ChatParser.NUMBER = 17;
ChatParser.WS = 18;

ChatParser.RULE_chat = 0;
ChatParser.RULE_line = 1;
ChatParser.RULE_message = 2;
ChatParser.RULE_name = 3;
ChatParser.RULE_command = 4;
ChatParser.RULE_verb = 5;
ChatParser.RULE_object = 6;
ChatParser.RULE_reason = 7;
ChatParser.RULE_method = 8;
ChatParser.RULE_time = 9;
ChatParser.RULE_emoticon = 10;
ChatParser.RULE_link = 11;
ChatParser.RULE_color = 12;
ChatParser.RULE_mention = 13;


function ChatContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_chat;
    return this;
}

ChatContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ChatContext.prototype.constructor = ChatContext;

ChatContext.prototype.EOF = function() {
    return this.getToken(ChatParser.EOF, 0);
};

ChatContext.prototype.line = function(i) {
    if(i===undefined) {
        i = null;
    }
    if(i===null) {
        return this.getTypedRuleContexts(LineContext);
    } else {
        return this.getTypedRuleContext(LineContext,i);
    }
};

ChatContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterChat(this);
	}
};

ChatContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitChat(this);
	}
};




ChatParser.ChatContext = ChatContext;

ChatParser.prototype.chat = function() {

    var localctx = new ChatContext(this, this._ctx, this.state);
    this.enterRule(localctx, 0, ChatParser.RULE_chat);
    var _la = 0; // Token type
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 29; 
        this._errHandler.sync(this);
        _la = this._input.LA(1);
        do {
            this.state = 28;
            this.line();
            this.state = 31; 
            this._errHandler.sync(this);
            _la = this._input.LA(1);
        } while(_la===ChatParser.IDENT);
        this.state = 33;
        this.match(ChatParser.EOF);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function LineContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_line;
    return this;
}

LineContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
LineContext.prototype.constructor = LineContext;

LineContext.prototype.name = function() {
    return this.getTypedRuleContext(NameContext,0);
};

LineContext.prototype.command = function() {
    return this.getTypedRuleContext(CommandContext,0);
};

LineContext.prototype.NEWLINE = function() {
    return this.getToken(ChatParser.NEWLINE, 0);
};

LineContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterLine(this);
	}
};

LineContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitLine(this);
	}
};




ChatParser.LineContext = LineContext;

ChatParser.prototype.line = function() {

    var localctx = new LineContext(this, this._ctx, this.state);
    this.enterRule(localctx, 2, ChatParser.RULE_line);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 35;
        this.name();
        this.state = 36;
        this.command();
        this.state = 37;
        this.match(ChatParser.NEWLINE);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function MessageContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_message;
    return this;
}

MessageContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
MessageContext.prototype.constructor = MessageContext;

MessageContext.prototype.emoticon = function(i) {
    if(i===undefined) {
        i = null;
    }
    if(i===null) {
        return this.getTypedRuleContexts(EmoticonContext);
    } else {
        return this.getTypedRuleContext(EmoticonContext,i);
    }
};

MessageContext.prototype.link = function(i) {
    if(i===undefined) {
        i = null;
    }
    if(i===null) {
        return this.getTypedRuleContexts(LinkContext);
    } else {
        return this.getTypedRuleContext(LinkContext,i);
    }
};

MessageContext.prototype.color = function(i) {
    if(i===undefined) {
        i = null;
    }
    if(i===null) {
        return this.getTypedRuleContexts(ColorContext);
    } else {
        return this.getTypedRuleContext(ColorContext,i);
    }
};

MessageContext.prototype.mention = function(i) {
    if(i===undefined) {
        i = null;
    }
    if(i===null) {
        return this.getTypedRuleContexts(MentionContext);
    } else {
        return this.getTypedRuleContext(MentionContext,i);
    }
};

MessageContext.prototype.IDENT = function(i) {
	if(i===undefined) {
		i = null;
	}
    if(i===null) {
        return this.getTokens(ChatParser.IDENT);
    } else {
        return this.getToken(ChatParser.IDENT, i);
    }
};


MessageContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterMessage(this);
	}
};

MessageContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitMessage(this);
	}
};




ChatParser.MessageContext = MessageContext;

ChatParser.prototype.message = function() {

    var localctx = new MessageContext(this, this._ctx, this.state);
    this.enterRule(localctx, 4, ChatParser.RULE_message);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 44; 
        this._errHandler.sync(this);
        var _alt = 1;
        do {
        	switch (_alt) {
        	case 1:
        		this.state = 44;
        		this._errHandler.sync(this);
        		switch(this._input.LA(1)) {
        		case ChatParser.T__0:
        		    this.state = 39;
        		    this.emoticon();
        		    break;
        		case ChatParser.TEXT:
        		    this.state = 40;
        		    this.link();
        		    break;
        		case ChatParser.T__8:
        		    this.state = 41;
        		    this.color();
        		    break;
        		case ChatParser.T__9:
        		    this.state = 42;
        		    this.mention();
        		    break;
        		case ChatParser.IDENT:
        		    this.state = 43;
        		    this.match(ChatParser.IDENT);
        		    break;
        		default:
        		    throw new antlr4.error.NoViableAltException(this);
        		}
        		break;
        	default:
        		throw new antlr4.error.NoViableAltException(this);
        	}
        	this.state = 46; 
        	this._errHandler.sync(this);
        	_alt = this._interp.adaptivePredict(this._input,2, this._ctx);
        } while ( _alt!=2 && _alt!=antlr4.atn.ATN.INVALID_ALT_NUMBER );
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function NameContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_name;
    return this;
}

NameContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
NameContext.prototype.constructor = NameContext;

NameContext.prototype.IDENT = function() {
    return this.getToken(ChatParser.IDENT, 0);
};

NameContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterName(this);
	}
};

NameContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitName(this);
	}
};




ChatParser.NameContext = NameContext;

ChatParser.prototype.name = function() {

    var localctx = new NameContext(this, this._ctx, this.state);
    this.enterRule(localctx, 6, ChatParser.RULE_name);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 48;
        this.match(ChatParser.IDENT);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function CommandContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_command;
    return this;
}

CommandContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
CommandContext.prototype.constructor = CommandContext;

CommandContext.prototype.message = function() {
    return this.getTypedRuleContext(MessageContext,0);
};

CommandContext.prototype.SAYS = function() {
    return this.getToken(ChatParser.SAYS, 0);
};

CommandContext.prototype.SHOUTS = function() {
    return this.getToken(ChatParser.SHOUTS, 0);
};

CommandContext.prototype.verb = function() {
    return this.getTypedRuleContext(VerbContext,0);
};

CommandContext.prototype.object = function() {
    return this.getTypedRuleContext(ObjectContext,0);
};

CommandContext.prototype.reason = function() {
    return this.getTypedRuleContext(ReasonContext,0);
};

CommandContext.prototype.method = function() {
    return this.getTypedRuleContext(MethodContext,0);
};

CommandContext.prototype.time = function() {
    return this.getTypedRuleContext(TimeContext,0);
};

CommandContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterCommand(this);
	}
};

CommandContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitCommand(this);
	}
};




ChatParser.CommandContext = CommandContext;

ChatParser.prototype.command = function() {

    var localctx = new CommandContext(this, this._ctx, this.state);
    this.enterRule(localctx, 8, ChatParser.RULE_command);
    var _la = 0; // Token type
    try {
        this.state = 72;
        this._errHandler.sync(this);
        switch(this._input.LA(1)) {
        case ChatParser.SAYS:
        case ChatParser.SHOUTS:
            this.enterOuterAlt(localctx, 1);
            this.state = 50;
            _la = this._input.LA(1);
            if(!(_la===ChatParser.SAYS || _la===ChatParser.SHOUTS)) {
            this._errHandler.recoverInline(this);
            }
            else {
            	this._errHandler.reportMatch(this);
                this.consume();
            }
            this.state = 51;
            this.match(ChatParser.T__0);
            this.state = 52;
            this.message();
            break;
        case ChatParser.IDENT:
            this.enterOuterAlt(localctx, 2);
            this.state = 53;
            this.verb();
            this.state = 55;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.T__1) {
                this.state = 54;
                this.match(ChatParser.T__1);
            }

            this.state = 58;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.IDENT) {
                this.state = 57;
                this.object();
            }

            this.state = 62;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.T__2) {
                this.state = 60;
                this.match(ChatParser.T__2);
                this.state = 61;
                this.reason();
            }

            this.state = 66;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.T__3) {
                this.state = 64;
                this.match(ChatParser.T__3);
                this.state = 65;
                this.method();
            }

            this.state = 70;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.T__4) {
                this.state = 68;
                this.match(ChatParser.T__4);
                this.state = 69;
                this.time();
            }

            break;
        default:
            throw new antlr4.error.NoViableAltException(this);
        }
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function VerbContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_verb;
    return this;
}

VerbContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
VerbContext.prototype.constructor = VerbContext;

VerbContext.prototype.IDENT = function() {
    return this.getToken(ChatParser.IDENT, 0);
};

VerbContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterVerb(this);
	}
};

VerbContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitVerb(this);
	}
};




ChatParser.VerbContext = VerbContext;

ChatParser.prototype.verb = function() {

    var localctx = new VerbContext(this, this._ctx, this.state);
    this.enterRule(localctx, 10, ChatParser.RULE_verb);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 74;
        this.match(ChatParser.IDENT);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function ObjectContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_object;
    return this;
}

ObjectContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ObjectContext.prototype.constructor = ObjectContext;

ObjectContext.prototype.IDENT = function() {
    return this.getToken(ChatParser.IDENT, 0);
};

ObjectContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterObject(this);
	}
};

ObjectContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitObject(this);
	}
};




ChatParser.ObjectContext = ObjectContext;

ChatParser.prototype.object = function() {

    var localctx = new ObjectContext(this, this._ctx, this.state);
    this.enterRule(localctx, 12, ChatParser.RULE_object);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 76;
        this.match(ChatParser.IDENT);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function ReasonContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_reason;
    return this;
}

ReasonContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ReasonContext.prototype.constructor = ReasonContext;

ReasonContext.prototype.message = function() {
    return this.getTypedRuleContext(MessageContext,0);
};

ReasonContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterReason(this);
	}
};

ReasonContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitReason(this);
	}
};




ChatParser.ReasonContext = ReasonContext;

ChatParser.prototype.reason = function() {

    var localctx = new ReasonContext(this, this._ctx, this.state);
    this.enterRule(localctx, 14, ChatParser.RULE_reason);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 78;
        this.message();
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function MethodContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_method;
    return this;
}

MethodContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
MethodContext.prototype.constructor = MethodContext;

MethodContext.prototype.message = function() {
    return this.getTypedRuleContext(MessageContext,0);
};

MethodContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterMethod(this);
	}
};

MethodContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitMethod(this);
	}
};




ChatParser.MethodContext = MethodContext;

ChatParser.prototype.method = function() {

    var localctx = new MethodContext(this, this._ctx, this.state);
    this.enterRule(localctx, 16, ChatParser.RULE_method);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 80;
        this.message();
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function TimeContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_time;
    return this;
}

TimeContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
TimeContext.prototype.constructor = TimeContext;

TimeContext.prototype.NUMBER = function() {
    return this.getToken(ChatParser.NUMBER, 0);
};

TimeContext.prototype.TIME_UNIT = function() {
    return this.getToken(ChatParser.TIME_UNIT, 0);
};

TimeContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterTime(this);
	}
};

TimeContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitTime(this);
	}
};




ChatParser.TimeContext = TimeContext;

ChatParser.prototype.time = function() {

    var localctx = new TimeContext(this, this._ctx, this.state);
    this.enterRule(localctx, 18, ChatParser.RULE_time);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 82;
        this.match(ChatParser.NUMBER);
        this.state = 83;
        this.match(ChatParser.TIME_UNIT);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function EmoticonContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_emoticon;
    return this;
}

EmoticonContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
EmoticonContext.prototype.constructor = EmoticonContext;


EmoticonContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterEmoticon(this);
	}
};

EmoticonContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitEmoticon(this);
	}
};




ChatParser.EmoticonContext = EmoticonContext;

ChatParser.prototype.emoticon = function() {

    var localctx = new EmoticonContext(this, this._ctx, this.state);
    this.enterRule(localctx, 20, ChatParser.RULE_emoticon);
    var _la = 0; // Token type
    try {
        this.state = 95;
        this._errHandler.sync(this);
        var la_ = this._interp.adaptivePredict(this._input,11,this._ctx);
        switch(la_) {
        case 1:
            this.enterOuterAlt(localctx, 1);
            this.state = 85;
            this.match(ChatParser.T__0);
            this.state = 87;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.T__5) {
                this.state = 86;
                this.match(ChatParser.T__5);
            }

            this.state = 89;
            this.match(ChatParser.T__6);
            break;

        case 2:
            this.enterOuterAlt(localctx, 2);
            this.state = 90;
            this.match(ChatParser.T__0);
            this.state = 92;
            this._errHandler.sync(this);
            _la = this._input.LA(1);
            if(_la===ChatParser.T__5) {
                this.state = 91;
                this.match(ChatParser.T__5);
            }

            this.state = 94;
            this.match(ChatParser.T__7);
            break;

        }
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function LinkContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_link;
    return this;
}

LinkContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
LinkContext.prototype.constructor = LinkContext;

LinkContext.prototype.TEXT = function(i) {
	if(i===undefined) {
		i = null;
	}
    if(i===null) {
        return this.getTokens(ChatParser.TEXT);
    } else {
        return this.getToken(ChatParser.TEXT, i);
    }
};


LinkContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterLink(this);
	}
};

LinkContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitLink(this);
	}
};




ChatParser.LinkContext = LinkContext;

ChatParser.prototype.link = function() {

    var localctx = new LinkContext(this, this._ctx, this.state);
    this.enterRule(localctx, 22, ChatParser.RULE_link);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 97;
        this.match(ChatParser.TEXT);
        this.state = 98;
        this.match(ChatParser.TEXT);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function ColorContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_color;
    return this;
}

ColorContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
ColorContext.prototype.constructor = ColorContext;

ColorContext.prototype.IDENT = function() {
    return this.getToken(ChatParser.IDENT, 0);
};

ColorContext.prototype.message = function() {
    return this.getTypedRuleContext(MessageContext,0);
};

ColorContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterColor(this);
	}
};

ColorContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitColor(this);
	}
};




ChatParser.ColorContext = ColorContext;

ChatParser.prototype.color = function() {

    var localctx = new ColorContext(this, this._ctx, this.state);
    this.enterRule(localctx, 24, ChatParser.RULE_color);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 100;
        this.match(ChatParser.T__8);
        this.state = 101;
        this.match(ChatParser.IDENT);
        this.state = 102;
        this.match(ChatParser.T__8);
        this.state = 103;
        this.message();
        this.state = 104;
        this.match(ChatParser.T__8);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


function MentionContext(parser, parent, invokingState) {
	if(parent===undefined) {
	    parent = null;
	}
	if(invokingState===undefined || invokingState===null) {
		invokingState = -1;
	}
	antlr4.ParserRuleContext.call(this, parent, invokingState);
    this.parser = parser;
    this.ruleIndex = ChatParser.RULE_mention;
    return this;
}

MentionContext.prototype = Object.create(antlr4.ParserRuleContext.prototype);
MentionContext.prototype.constructor = MentionContext;

MentionContext.prototype.IDENT = function() {
    return this.getToken(ChatParser.IDENT, 0);
};

MentionContext.prototype.enterRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.enterMention(this);
	}
};

MentionContext.prototype.exitRule = function(listener) {
    if(listener instanceof ChatListener ) {
        listener.exitMention(this);
	}
};




ChatParser.MentionContext = MentionContext;

ChatParser.prototype.mention = function() {

    var localctx = new MentionContext(this, this._ctx, this.state);
    this.enterRule(localctx, 26, ChatParser.RULE_mention);
    try {
        this.enterOuterAlt(localctx, 1);
        this.state = 106;
        this.match(ChatParser.T__9);
        this.state = 107;
        this.match(ChatParser.IDENT);
    } catch (re) {
    	if(re instanceof antlr4.error.RecognitionException) {
	        localctx.exception = re;
	        this._errHandler.reportError(this, re);
	        this._errHandler.recover(this, re);
	    } else {
	    	throw re;
	    }
    } finally {
        this.exitRule();
    }
    return localctx;
};


exports.ChatParser = ChatParser;
