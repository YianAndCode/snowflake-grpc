<?php
require_once __DIR__ . "/vendor/autoload.php";

use Snowflake\SnowflakeClient;
use Snowflake\NextRequest;
use Snowflake\ParseRequest;

$client = new SnowflakeClient(
    "127.0.0.1:6666",
    [
        'credentials' => Grpc\ChannelCredentials::createInsecure()
    ]
);

$request = new NextRequest();
$request->setServiceId(mt_rand(0, 31));

$get = $client->Next($request)->wait();
list($reply, $status) = $get;

if ($status->code != 0)
{
    echo "Call Next() failed, err code: {$status->code}." . PHP_EOL;
    return;
}

$id = $reply->getId();
echo $id . PHP_EOL;

$request = new ParseRequest;
$request->setId($id);
$get = $client->Parse($request)->wait();
list($reply, $status) = $get;
if ($status->code != 0)
{
    echo "Call Parse() failed, err code: {$status->code}." . PHP_EOL;
    return;
}
$data['timestamp'] = $reply->getTimestamp();
$data['nodeID'] = $reply->getNodeId();
$data['serviceID'] = $reply->getServiceId();
$data['seq'] = $reply->getSeq();
print_r($data);