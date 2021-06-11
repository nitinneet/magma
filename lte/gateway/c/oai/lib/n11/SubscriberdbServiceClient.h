/**
 * Copyright 2020 The Magma Authors.
 *
 * This source code is licensed under the BSD-style license found in the
 * LICENSE file in the root directory of this source tree.
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
#pragma once

#include <grpc++/grpc++.h>
#include "lte/protos/subscriberauth.grpc.pb.h"
#include "lte/protos/subscriberauth.pb.h"
#include "GRPCReceiver.h"
#include <stdint.h>
#include <functional>
#include <memory>
#include "3gpp_38.413.h"

using grpc::Status;
using magma::GRPCReceiver;
using magma::lte::M5GAuthenticationInformationAnswer;
using magma::lte::M5GAuthenticationInformationRequest;
using magma::lte::M5GSubscriberAuthentication;

namespace magma5g {

class SubscriberdbServiceClient {
 public:
  virtual ~SubscriberdbServiceClient() {}
  virtual bool get_subs_auth_info(
      const std::string& imsi, uint8_t imsi_length, amf_ue_ngap_id_t ue_id) = 0;

  virtual bool get_subs_auth_info_resync(
      const std::string& imsi, uint8_t imsi_length, const void* resync_info,
      uint8_t resync_info_len, amf_ue_ngap_id_t ue_id) = 0;
};

/**
 * AsyncSmfServiceClient implements SmfServiceClient but sends calls
 * asynchronously to sessiond.
 */
class AsyncSubscriberdbServiceClient : public GRPCReceiver,
                                       public SubscriberdbServiceClient {
 public:
  bool get_subs_auth_info(
      const std::string& imsi, uint8_t imsi_length, amf_ue_ngap_id_t ue_id);

  bool get_subs_auth_info_resync(
      const std::string& imsi, uint8_t imsi_length, const void* resync_info,
      uint8_t resync_info_len, amf_ue_ngap_id_t ue_id);

  static AsyncSubscriberdbServiceClient& getInstance();

  AsyncSubscriberdbServiceClient(AsyncSubscriberdbServiceClient const&) =
      delete;
  void operator=(AsyncSubscriberdbServiceClient const&) = delete;

 private:
  AsyncSubscriberdbServiceClient();
  static const uint32_t RESPONSE_TIMEOUT = 10;  // seconds
  std::unique_ptr<M5GSubscriberAuthentication::Stub> stub_{};

  void GetSubscriberAuthInfoRPC(
      M5GAuthenticationInformationRequest& request,
      const std::function<
          void(grpc::Status, M5GAuthenticationInformationAnswer)>& callback);
};
}  // namespace magma5g
