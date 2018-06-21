<?php
namespace Snowflake;

class SnowflakeClient extends \Grpc\BaseStub
{
    public function __construct($host, $opts, $channel = null)
    {
        parent::__construct($host, $opts, $channel);
    }

    public function Next(NextRequest $argument, $metadata = [], $options = [])
    {
        return $this->_simpleRequest(
            '/snowflake.Snowflake/Next',
            $argument,
            ['\Snowflake\NextReply', 'decode'],
            $metadata,
            $options
        );
    }

    public function Parse(ParseRequest $argument, $metadata = [], $options = [])
    {
        return $this->_simpleRequest(
            '/snowflake.Snowflake/Parse',
            $argument,
            ['\Snowflake\ParseReply', 'decode'],
            $metadata,
            $options
        );
    }
}